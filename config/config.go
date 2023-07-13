package config

import (
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
)

func Load() {
	_ = log.Load()
	_ = web.Load()
	_ = db.Load()
	_ = mqtt.Load()
}
