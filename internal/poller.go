package internal

import (
	"context"
	"fmt"
	"github.com/iot-master-contrib/gateway/connect"
	"github.com/iot-master-contrib/gateway/protocols"
	"github.com/iot-master-contrib/gateway/protocols/define"
	"github.com/iot-master-contrib/gateway/types"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"time"
)

type poller struct {
	adapter define.Adapter
	devices []*types.Device
}

func newPoller(tunnel connect.Conn, options types.ProtocolOptions) (*poller, error) {
	adapter, err := protocols.Create(tunnel, options.Name, options.Options)
	if err != nil {
		return nil, err
	}
	return &poller{adapter: adapter}, nil
}

func (p *poller) Load(tunnel string) error {
	return db.Engine.Where("tunnel_id=?", tunnel).Find(&p.devices)
}

func (p *poller) Poll() bool {
	total := 0

	//TODO 将迭代器提升到p中，单次调用只查询一个设备
	for _, device := range p.devices {
		product := Products.Load(device.ProductId)
		if product == nil {
			continue
		}

		//统计加1
		sum := 0

		values, err := p.adapter.Poll(device.Station, product.Mappers)
		if err != nil {
			log.Error(err)
			//continue
		}

		if len(values) > 0 {
			total += sum

			//过滤字段
			for i, c := range product.filters {
				ret, err := c.EvalBool(context.Background(), values)
				if err != nil {
					log.Error(err)
					continue
				}
				if !ret {
					name := product.Filters[i].Name
					if name == "*" {
						//break xx
						//TODO 不上传数据
					} else {
						delete(values, name)
					}
				}
			}

			//计算数据
			for i, c := range product.calculators {
				ret, err := c(context.Background(), values)
				if err != nil {
					log.Error(err)
					continue
				}
				name := product.Calculators[i].Name
				values[name] = ret
			}

			//mqtt上传数据，暂定使用Object方式，简单
			topic := fmt.Sprintf("up/property/%s/%s", product.Id, device.Id)
			_ = mqtt.Publish(topic, values)

			//上线提醒
			if !device.Online {
				device.Online = true
				Online(device.ProductId, device.Id)
			}
		} else {
			//掉线提醒
			if device.Online {
				device.Online = false
				Offline(device.ProductId, device.Id)
			}
		}
	}

	//如果没有设备，就睡眠1分钟
	if total == 0 {
		time.Sleep(time.Second * 5)
		//return errors.New("没有设备")
	}

	return true
}

func (p *poller) Close() error {

	for _, device := range p.devices {
		Offline(device.ProductId, device.Id)
	}

	return nil
}
