package admin

import (
	"github.com/zgwit/iot-gateway/api"
)

func init() {
	api.Register("GET", "logout", logout)
	api.Register("POST", "password", password)
}
