package connect

import (
	"errors"
	"github.com/zgwit/iot-master/v2/model"
	"net"
)

// ServerTcpTunnel 网络连接
type ServerTcpTunnel struct {
	tunnelBase
}

func newServerTcpTunnel(tunnel *model.Tunnel, conn net.Conn) *ServerTcpTunnel {
	return &ServerTcpTunnel{tunnelBase: tunnelBase{
		tunnel: tunnel,
		link:   conn,
	}}
}

func (l *ServerTcpTunnel) Open() error {
	return errors.New("ServerTcpTunnel cannot open")
}
