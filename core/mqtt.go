package core

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/timshannon/bolthold"
	"github.com/zgwit/iot-master/v2/model"
	"github.com/zgwit/iot-master/v2/pkg/log"
	"iot-master-gateway/db"
	"iot-master-gateway/dbus"
)

func RegisterGatewayHandler(client mqtt.Client, id string) {
	//id := config.Config.Node
	download := "/gateway/" + id + "/download/"

	client.Subscribe(download+"#", 0, func(client mqtt.Client, message mqtt.Message) {
		log.Infof("download %s", message.Topic())
	})

	client.AddRoute(download+"tunnel", func(client mqtt.Client, message mqtt.Message) {
		var tunnel model.Tunnel
		err := json.Unmarshal(message.Payload(), &tunnel)
		if err != nil {
			log.Error(err)
			return
		}

		err = db.Store().Upsert(tunnel.Id, &tunnel)
		if err != nil {
			log.Error(err)
			return
		}

		//TODO start tunnel
	})

	client.AddRoute(download+"server", func(client mqtt.Client, message mqtt.Message) {
		var server model.Server
		err := json.Unmarshal(message.Payload(), &server)
		if err != nil {
			log.Error(err)
			return
		}

		err = db.Store().Upsert(server.Id, &server)
		if err != nil {
			log.Error(err)
			return
		}

		//TODO start server

	})

	client.AddRoute(download+"product", func(client mqtt.Client, message mqtt.Message) {
		var product model.Product
		err := json.Unmarshal(message.Payload(), &product)
		if err != nil {
			log.Error(err)
			return
		}

		err = db.Store().Upsert(product.Id, &product)
		if err != nil {
			log.Error(err)
			return
		}

		//TODO start devices

	})

	client.AddRoute(download+"device", func(client mqtt.Client, message mqtt.Message) {
		var device model.Device
		err := json.Unmarshal(message.Payload(), &device)
		if err != nil {
			log.Error(err)
			return
		}

		err = db.Store().Insert(device.Id, &device)
		if err != nil {
			log.Error(err)
			return
		}

		//检查产品
		var product model.Product
		err = db.Store().Get(device.ProductId, &product)
		if err != nil {
			log.Error(err)
			if err == bolthold.ErrNotFound {
				err = dbus.Publish("/gateway/"+id+"/command/product/"+device.ProductId, []byte(""))
				if err != nil {
					log.Error(err)
				}
			}
			return
		}

		//TODO start device
		var devices []model.Device
		err = db.Store().Find(&devices, bolthold.Where("ProductId").Eq(device.ProductId))
		if err != nil {
			log.Error(err)
			return
		}
		//for _, dev := range devices {
		//	//
		//}

	})

	client.AddRoute(download+"config", func(client mqtt.Client, message mqtt.Message) {

	})

	command := "/gateway/" + id + "/command/"

	client.Subscribe(command+"#", 0, func(client mqtt.Client, message mqtt.Message) {
		log.Infof("command %s", message.Topic())
	})
	client.AddRoute(command+"start", func(client mqtt.Client, message mqtt.Message) {

	})
	client.AddRoute(command+"stop", func(client mqtt.Client, message mqtt.Message) {

	})
	client.AddRoute(command+"restart", func(client mqtt.Client, message mqtt.Message) {

	})
	client.AddRoute(command+"status", func(client mqtt.Client, message mqtt.Message) {

	})

}

func RegisterTunnelHandler(client mqtt.Client, id string) {
	command := "/tunnel/" + id + "/command/"

	client.Subscribe(command+"#", 0, func(client mqtt.Client, message mqtt.Message) {
		log.Infof("tunnel command %s", message.Topic())
	})
	client.AddRoute(command+"start", func(client mqtt.Client, message mqtt.Message) {

	})
	client.AddRoute(command+"stop", func(client mqtt.Client, message mqtt.Message) {

	})
	client.AddRoute(command+"restart", func(client mqtt.Client, message mqtt.Message) {

	})
	client.AddRoute(command+"status", func(client mqtt.Client, message mqtt.Message) {

	})
}

func RegisterServerHandler(client mqtt.Client, id string) {
	command := "/server/" + id + "/command/"

	client.Subscribe(command+"#", 0, func(client mqtt.Client, message mqtt.Message) {
		log.Infof("server command %s", message.Topic())
	})
	client.AddRoute(command+"start", func(client mqtt.Client, message mqtt.Message) {

	})
	client.AddRoute(command+"stop", func(client mqtt.Client, message mqtt.Message) {

	})
	client.AddRoute(command+"restart", func(client mqtt.Client, message mqtt.Message) {

	})
	client.AddRoute(command+"status", func(client mqtt.Client, message mqtt.Message) {

	})
}

func RegisterDeviceHandler(client mqtt.Client, id string) {
	command := "/device/" + id + "/command/"

	client.Subscribe(command+"#", 0, func(client mqtt.Client, message mqtt.Message) {
		log.Infof("tunnel command %s", message.Topic())
	})
	client.AddRoute(command+"start", func(client mqtt.Client, message mqtt.Message) {

	})
	client.AddRoute(command+"stop", func(client mqtt.Client, message mqtt.Message) {

	})
	client.AddRoute(command+"restart", func(client mqtt.Client, message mqtt.Message) {

	})
	client.AddRoute(command+"status", func(client mqtt.Client, message mqtt.Message) {

	})
}
