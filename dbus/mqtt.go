package dbus

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"iot-master-gateway/config"
)

func Open() error {
	opts := mqtt.NewClientOptions()
	client := mqtt.NewClient(opts)

	topic := fmt.Sprintf("$GW/%s/s", config.Config.Node)
	client.Subscribe(topic, 0, func(client mqtt.Client, message mqtt.Message) {

	})
}
