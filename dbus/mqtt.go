package dbus

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgwit/iot-master/v2/pkg/log"
)

var MQTT mqtt.Client

func Open(cfg Options) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.Broker)
	opts.SetClientID(cfg.ClientID)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	opts.SetConnectRetry(true)
	//opts.SetMaxReconnectInterval(time.Minute)
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		log.Info("dbus connected")
	})
	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Info("dbus lost", err.Error())
	}

	MQTT = mqtt.NewClient(opts)
}

func Publish(topic string, payload interface{}) (err error) {
	switch payload.(type) {
	case nil:
		payload = []byte("")
	case struct{}:
		payload, err = json.Marshal(payload)
		if err != nil {
			return
		}
	}
	token := MQTT.Publish(topic, 0, false, payload)
	return token.Error()
}
