package modbus

import (
	"errors"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/convert"
	"github.com/god-jason/bucket/pool"
	"github.com/god-jason/bucket/types"
	"github.com/zgwit/iot-gateway/connect"
	"github.com/zgwit/iot-gateway/device"
	"github.com/zgwit/iot-gateway/product"
	"slices"
	"time"
)

type Adapter struct {
	tunnel connect.Tunnel
	modbus Modbus

	devices []*device.Device
	index   map[string]*device.Device

	//product_id => xxx
	mappers map[string]*Mapper
	pollers map[string]*[]*Poller

	options types.Options
	//index map[string]*device.Device
}

func (adapter *Adapter) Tunnel() connect.Tunnel {
	return adapter.tunnel
}

func (adapter *Adapter) start() (err error) {
	adapter.devices, err = device.LoadByTunnel(adapter.tunnel.ID())
	if err != nil {
		return err
	}

	//if len(adapter.devices) == 0 {
	//	return errors.New("无设备")
	//}

	for _, d := range adapter.devices {
		//索引
		adapter.index[d.Id] = d

		//加载映射表
		adapter.mappers[d.ProductId], err = product.LoadConfig[Mapper](d.ProductId, "mapper")
		if err != nil {
			log.Error(err)
		}

		//加载轮询表
		adapter.pollers[d.ProductId], err = product.LoadConfig[[]*Poller](d.ProductId, "poller")
		if err != nil {
			log.Error(err)
		}
	}

	//开始轮询
	go adapter.poll()
	return nil
}

func (adapter *Adapter) poll() {

	//设备上线
	//!!! 不能这样做，不然启动服务器会产生大量的消息
	//for _, dev := range adapter.index {
	//	topic := fmt.Sprintf("device/online/%s", dev.Id)
	//	_ = mqtt.Publish(topic, nil)
	//}

	interval := adapter.options.Int64("poller_interval", 60) //默认1分钟轮询一次
	if interval < 1 {
		interval = 1
	}

	//按毫秒计时
	interval *= 1000

	//OUT:
	for {
		start := time.Now().UnixMilli()
		for _, dev := range adapter.devices {
			values, err := adapter.Sync(dev.Id)
			if err != nil {
				log.Error(err)
				continue
			}

			//d := device.Get(dev.Id)
			if values != nil && len(values) > 0 {
				_ = pool.Insert(func() {
					dev.Push(values)
				})
			}
		}

		//检查连接，避免空等待
		if !adapter.tunnel.Available() {
			break
		}

		//轮询间隔
		now := time.Now().UnixMilli()
		elapsed := now - start
		if elapsed < interval {
			time.Sleep(time.Millisecond * time.Duration(interval-elapsed))
		}

		//避免空转，睡眠1分钟（延迟10ms太长，睡1分钟也有点长）
		if elapsed < 10 {
			time.Sleep(time.Minute)
		}
	}

	log.Info("modbus adapter quit", adapter.tunnel.ID())

	//设备下线
	//for _, dev := range adapter.devices {
	//	topic := fmt.Sprintf("device/%s/offline", dev.Id)
	//	_ = mqtt.Publish(topic, nil)
	//}

	//TODO d.SetAdapter(nil)
}

func (adapter *Adapter) Mount(id string) error {
	dev, err := device.Ensure(id)
	if err != nil {
		return err
	}

	found := false
	for i, d := range adapter.devices {
		if d.Id == id {
			adapter.devices[i] = dev
			adapter.index[id] = dev
			found = true
		}
	}
	if !found {
		adapter.devices = append(adapter.devices, dev)
		adapter.index[id] = dev
	}
	return nil
}

func (adapter *Adapter) Unmount(id string) error {
	delete(adapter.index, id)
	for i, d := range adapter.devices {
		if d.Id == id {
			slices.Delete(adapter.devices, i, i+1)
			return nil
		}
	}
	return nil
}

func (adapter *Adapter) Get(id, name string) (any, error) {
	d := adapter.index[id]
	station := d.Station["slave"]

	//todo error
	mapper, code, address := adapter.mappers[d.ProductId].Lookup(name)
	if mapper == nil {
		return nil, errors.New("找不到数据点")
	}

	//此处全部读取了，有些冗余
	data, err := adapter.modbus.Read(convert.ToUint8(station), code, address, 2)
	if err != nil {
		return nil, err
	}

	return mapper.Parse(address, data)
}

func (adapter *Adapter) Set(id, name string, value any) error {
	d := adapter.index[id]
	station := d.Station["slave"]

	mapper, code, address := adapter.mappers[d.ProductId].Lookup(name)
	if mapper == nil {
		return errors.New("地址找不到")
	}

	data, err := mapper.Encode(value)
	if err != nil {
		return err
	}
	return adapter.modbus.Write(convert.ToUint8(station), code, address, data)
}

func (adapter *Adapter) Sync(id string) (map[string]any, error) {
	d := adapter.index[id]
	station := convert.ToUint8(d.Station["slave"])

	//没有地址表和轮询器，则跳过
	//if d.pollers == nil || d.mappers == nil {
	//	return nil, nil
	//}

	values := make(map[string]any)
	for _, poller := range *adapter.pollers[d.ProductId] {
		data, err := adapter.modbus.Read(station, poller.Code, poller.Address, poller.Length)
		if err != nil {
			return nil, err
		}
		err = poller.Parse(adapter.mappers[d.ProductId], data, values)
		if err != nil {
			return nil, err
		}
	}

	//TODO 计算器

	return values, nil
}
