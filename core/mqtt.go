package core

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v2/pkg/log"
)

func RegisterGatewayHandler(client mqtt.Client, id string) {
	//id := config.Config.Node
	download := "/gateway/" + id + "/download/"

	client.Subscribe(download+"#", 0, func(client mqtt.Client, message mqtt.Message) {
		log.Infof("download %s", message.Topic())
	})

	client.AddRoute(download+"tunnel", func(client mqtt.Client, message mqtt.Message) {

	})

	client.AddRoute(download+"server", func(client mqtt.Client, message mqtt.Message) {

	})

	client.AddRoute(download+"product", func(client mqtt.Client, message mqtt.Message) {

	})

	client.AddRoute(download+"device", func(client mqtt.Client, message mqtt.Message) {

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
