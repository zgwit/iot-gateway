package modbus

import (
	"errors"
	"github.com/god-jason/bucket/types"
	"github.com/zgwit/iot-gateway/product"
)

type Adapter struct {
	modbus Modbus

	//device=>product_id
	devices  map[string]string
	stations map[string]types.Options

	//product_id => xxx
	mappers map[string]*Mapper
	pollers map[string]*[]*Poller
}

func (adapter *Adapter) Mount(id string, product_id string, station types.Options) (err error) {
	adapter.devices[id] = product_id
	adapter.stations[id] = station

	//加载映射表
	adapter.mappers[product_id], err = product.LoadConfig[Mapper](product_id, "mapper")
	if err != nil {
		return err
	}

	//加载轮询表
	adapter.pollers[product_id], err = product.LoadConfig[[]*Poller](product_id, "poller")
	if err != nil {
		return err
	}

	return nil
}

func (adapter *Adapter) Unmount(id string) error {
	delete(adapter.devices, id)
	delete(adapter.stations, id)
	return nil
}

func (adapter *Adapter) Get(id, name string) (any, error) {
	product_id := adapter.devices[id]
	station := adapter.stations[id]
	slave := station.Int("slave", 1)

	mapper := adapter.mappers[product_id]
	if mapper == nil {
		return nil, errors.New("没有地址映射")
	}
	point, code, address := mapper.Lookup(name)
	if point == nil {
		return nil, errors.New("找不到数据点")
	}

	//此处全部读取了，有些冗余
	data, err := adapter.modbus.Read(uint8(slave), code, address, 2)
	if err != nil {
		return nil, err
	}

	return point.Parse(address, data)
}

func (adapter *Adapter) Set(id, name string, value any) error {
	product_id := adapter.devices[id]
	station := adapter.stations[id]
	slave := station.Int("slave", 1)

	mapper := adapter.mappers[product_id]
	if mapper == nil {
		return errors.New("没有地址映射")
	}
	point, code, address := mapper.Lookup(name)
	if point == nil {
		return errors.New("地址找不到")
	}

	data, err := point.Encode(value)
	if err != nil {
		return err
	}
	return adapter.modbus.Write(uint8(slave), code, address, data)
}

func (adapter *Adapter) Sync(id string) (map[string]any, error) {
	product_id := adapter.devices[id]
	station := adapter.stations[id]
	slave := station.Int("slave", 1)

	//没有地址表和轮询器，则跳过
	//if d.pollers == nil || d.mappers == nil {
	//	return nil, nil
	//}
	mapper := adapter.mappers[product_id]
	if mapper == nil {
		return nil, errors.New("没有地址映射")
	}

	values := make(map[string]any)
	for _, poller := range *adapter.pollers[product_id] {
		if poller == nil {
			continue
		}
		data, err := adapter.modbus.Read(uint8(slave), poller.Code, poller.Address, poller.Length)
		if err != nil {
			return nil, err
		}
		err = poller.Parse(mapper, data, values)
		if err != nil {
			return nil, err
		}
	}

	//TODO 过滤器

	//TODO 计算器

	return values, nil
}
