package core

import (
	"github.com/zgwit/go-plc/protocol"
	"github.com/zgwit/iot-master/v2/pkg/log"
	"iot-master-gateway/connect"
	"iot-master-gateway/dbus"
	"time"
)

type Tunnel struct {
	tunnel  connect.Tunnel
	adapter protocol.Protocol

	devices []*Device //lock
}

func (t *Tunnel) Poll() {
	//TODO 此处解析地址？

	for {
		start := time.Now().UnixMilli()

		for _, device := range t.devices {
			//device.product.Pollers
			for _, poller := range device.product.pollers {
				buf, err := t.adapter.Read(device.Station, poller.addr, poller.Length)
				if err != nil {
					log.Error(err)
					continue
				}
				values, err := device.product.Parse(buf, poller.addr)
				if err != nil {
					log.Error(err)
					continue
				}

				err = dbus.Publish("/device/"+device.Id+"/values", values)
				if err != nil {
					log.Error(err)
				}
			}

			//TODO 在此处统一上报
		}

		end := time.Now().UnixMilli()

		//轮询间隔时间，计算
		remain := 60000 - end + start
		if remain > 0 {
			time.Sleep(time.Millisecond * time.Duration(remain))
		}
	}
}
