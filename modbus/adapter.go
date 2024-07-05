package modbus

import (
	"errors"
	"fmt"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pool"
	"github.com/zgwit/iot-gateway/connect"
	"github.com/zgwit/iot-gateway/db"
	"github.com/zgwit/iot-gateway/mqtt"
	"github.com/zgwit/iot-gateway/product"
	"github.com/zgwit/iot-gateway/types"
	"slices"
	"time"
)

type Adapter struct {
	tunnel  connect.Tunnel
	modbus  Modbus
	devices []*Device

	index map[string]*Device

	options types.Options
	//index map[string]*device.Device
}

func (adapter *Adapter) Tunnel() connect.Tunnel {
	return adapter.tunnel
}

func (adapter *Adapter) start() error {
	err := db.Engine.Where("tunnel_id=?", adapter.tunnel.ID()).
		And("disabled!=1").Find(&adapter.devices)
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
		d.mapper, err = product.LoadConfig[Mapper](d.ProductId, "mapper")
		if err != nil {
			log.Error(err)
		}

		//加载轮询表
		d.pollers, err = product.LoadConfig[[]*Poller](d.ProductId, "poller")
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
					topic := fmt.Sprintf("device/"+dev.Id+"/values", dev.Id)
					mqtt.Publish(topic, values)
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

func (adapter *Adapter) Mount(device string) error {
	var dev Device
	has, err := db.Engine.ID(device).Get(&dev)
	if err != nil {
		return err
	}
	if !has {
		return errors.New("找不到设备")
	}

	found := false
	for i, d := range adapter.devices {
		if d.Id == device {
			adapter.devices[i] = &dev
			adapter.index[device] = &dev
			found = true
		}
	}
	if !found {
		adapter.devices = append(adapter.devices, &dev)
		adapter.index[device] = &dev
	}
	return nil
}

func (adapter *Adapter) Unmount(device string) error {
	delete(adapter.index, device)
	for i, d := range adapter.devices {
		if d.Id == device {
			slices.Delete(adapter.devices, i, i+1)
			return nil
		}
	}
	return nil
}

func (adapter *Adapter) Get(id, name string) (any, error) {
	d := adapter.index[id]
	station := d.Station.Slave

	mapper, code, address := d.mapper.Lookup(name)
	if mapper == nil {
		return nil, errors.New("找不到数据点")
	}

	//此处全部读取了，有些冗余
	data, err := adapter.modbus.Read(station, code, address, 2)
	if err != nil {
		return nil, err
	}

	return mapper.Parse(address, data)
}

func (adapter *Adapter) Set(id, name string, value any) error {
	d := adapter.index[id]
	station := d.Station.Slave

	mapper, code, address := d.mapper.Lookup(name)
	if mapper == nil {
		return errors.New("地址找不到")
	}

	data, err := mapper.Encode(value)
	if err != nil {
		return err
	}
	return adapter.modbus.Write(station, code, address, data)
}

func (adapter *Adapter) Sync(id string) (map[string]any, error) {
	d := adapter.index[id]
	station := d.Station.Slave

	//没有地址表和轮询器，则跳过
	if d.pollers == nil || d.mapper == nil {
		return nil, nil
	}

	values := make(map[string]any)
	for _, poller := range *d.pollers {
		data, err := adapter.modbus.Read(station, poller.Code, poller.Address, poller.Length)
		if err != nil {
			return nil, err
		}
		err = poller.Parse(d.mapper, data, values)
		if err != nil {
			return nil, err
		}
	}

	//TODO 计算器

	return values, nil
}
