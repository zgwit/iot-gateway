package server

import (
	"errors"
	"github.com/zgwit/iot-gateway/db"
	"github.com/zgwit/iot-gateway/gateway"
)

func init() {
	db.Register(new(Link))
}

// Link 网络连接
type Link struct {
	gateway.Base `xorm:"extends"`

	ServerId string `json:"server_id" xorm:"index"` //服务器ID
	Remote   string `json:"remote,omitempty"`       //远程地址
}

func (l *Link) Open() error {

	return errors.New("link cannot open")
}
