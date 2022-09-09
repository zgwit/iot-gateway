package core

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"iot-master-gateway/dbus"
	"iot-master-gateway/log"
)

func RegisterHandler() {

	// gateway/{gid}/download/product/list/{tid...}
	// gateway/{gid}/response/product/list/{tid...}

	download := "gateway/{id}/download/"
	//response := "gateway/{id}/response/{tid}"
	dbus.MQTT.Subscribe(download+"#", 0, func(client mqtt.Client, message mqtt.Message) {
		log.Infof("dbus api %s", message.Topic())
	})

	dbus.MQTT.AddRoute(download+"tunnel", func(client mqtt.Client, message mqtt.Message) {

	})

	dbus.MQTT.AddRoute(download+"server", func(client mqtt.Client, message mqtt.Message) {

	})

	dbus.MQTT.AddRoute(download+"product", func(client mqtt.Client, message mqtt.Message) {

	})

	dbus.MQTT.AddRoute(download+"device", func(client mqtt.Client, message mqtt.Message) {

	})

	dbus.MQTT.AddRoute(download+"config", func(client mqtt.Client, message mqtt.Message) {

	})

}
