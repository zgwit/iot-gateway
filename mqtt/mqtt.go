package mqtt

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"iot-master-gateway/config"
	"iot-master-gateway/log"
)

var MQTT mqtt.Client

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

type command struct {
	Id   string
	Body []byte
}

func test() {

	request := "gateway/{id}/request/"
	response := "gateway/{id}/response/"
	MQTT.Subscribe(request+"#", 0, func(client mqtt.Client, message mqtt.Message) {
		log.Infof("mqtt api %s", message.Topic())
		//MQTT.Publish()
	})

	MQTT.AddRoute(request+"product", func(client mqtt.Client, message mqtt.Message) {

		_ = Publish(response+"product", "")
	})
	MQTT.AddRoute(request+"product/list", func(client mqtt.Client, message mqtt.Message) {

	})

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
