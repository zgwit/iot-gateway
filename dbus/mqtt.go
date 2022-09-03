package dbus

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"iot-master-gateway/config"
	"iot-master-gateway/log"
)

var MQTT mqtt.Client
var Master mqtt.Client

func Open(cfg config.MQTT) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.Broker)
	opts.SetClientID(cfg.ClientID)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	opts.SetConnectRetry(true)
	//opts.SetMaxReconnectInterval(time.Minute)
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		log.Info("mqtt connected")
	})
	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Info("mqtt lost", err.Error())
	}

	MQTT = mqtt.NewClient(opts)
}

func OpenMaster(cfg config.MQTT) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.Broker)
	opts.SetClientID(cfg.ClientID)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	opts.SetConnectRetry(true)
	//opts.SetMaxReconnectInterval(time.Minute)
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		log.Info("master connected")
	})
	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Info("master lost", err.Error())
	}

	Master = mqtt.NewClient(opts)

	//topic := fmt.Sprintf("$GW/%s/s", config.Config.Node)
	//Master.Subscribe(topic, 0, func(client mqtt.Client, message mqtt.Message) {
}
