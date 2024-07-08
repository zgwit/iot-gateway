package device

import (
	"encoding/json"
	"github.com/god-jason/bucket/log"
	"github.com/zgwit/iot-gateway/mqtt"
	"strings"
)

func OnValuesChange() {
	//数值
	mqtt.Subscribe("down/device/+/values", func(topic string, payload []byte) {
		var values map[string]any
		err := json.Unmarshal(payload, &values)
		if err != nil {
			log.Error(err)
			return
		}
		ss := strings.Split(topic, "/")
		id := ss[2]
		dev := devices.Load(id)
		if dev == nil {
			err := dev.WriteMany(values)
			if err != nil {
				log.Error(err)
			}
		}
	})

	mqtt.Subscribe("down/device/+/action", func(topic string, payload []byte) {
		var values map[string]any
		err := json.Unmarshal(payload, &values)
		if err != nil {
			log.Error(err)
			return
		}
		ss := strings.Split(topic, "/")
		id := ss[2]
		dev := devices.Load(id)
		if dev == nil {
			//todo 执行动作
			//err := dev.Action(values)

		}
	})

}
