package core

import (
	"github.com/zgwit/go-plc/protocol"
	"github.com/zgwit/iot-master/v2/pkg/log"
	"iot-master-gateway/connect"
	"iot-master-gateway/dbus"
)

type Tunnel struct {
	tunnel  connect.Tunnel
	adapter protocol.Protocol

	devices []*Device //lock
}

func (t *Tunnel) Poll() {
	for {
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
		}
	}
}
