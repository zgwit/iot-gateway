package core

import (
	"github.com/zgwit/go-plc/protocol"
	"iot-master-gateway/connect"
)

type Tunnel struct {
	tunnel  connect.Tunnel
	adapter protocol.Protocol

	devices []*Device //lock
}
