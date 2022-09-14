package core

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

func (l *ServerTcpTunnel) receive() {
	l.running = true
	l.online = true

	buf := make([]byte, 1024)
	for {
		n, err := l.link.Read(buf)
		if err != nil {
			l.onClose()
			break
		}
		if n == 0 {
			continue
		}
		//透传转发
		if l.pipe != nil {
			_, err = l.pipe.Write(buf[:n])
			if err != nil {
				l.pipe = nil
			} else {
				continue
			}
		}
	}
	l.running = false
	l.online = false
}
