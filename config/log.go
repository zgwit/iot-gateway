package config

import "github.com/zgwit/iot-master/v2/pkg/log"

var LogDefault = log.Options{
	Level:  "trace",
	Caller: true,
	Text:   false,
}
