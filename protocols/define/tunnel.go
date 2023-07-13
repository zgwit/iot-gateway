package define

import "github.com/iot-master-contrib/gateway/connect"

// Tunnel 通道
type Tunnel interface {
	connect.Conn

	Open() error

	Running() bool

	Online() bool
}
