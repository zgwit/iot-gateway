package connect

import (
	"fmt"
	"github.com/zgwit/iot-master/v2/model"
	"io"
	"strings"
)

// Tunnel 通道
type Tunnel interface {
	io.ReadWriteCloser

	Model() *model.Tunnel

	Open() error

	Running() bool

	Online() bool
}

// NewTunnel 创建通道
func NewTunnel(tunnel *model.Tunnel) (Tunnel, error) {
	var tnl Tunnel
	switch tunnel.Type {
	case "serial":
		tnl = newTunnelSerial(tunnel)
		break
	case "tcp-client":
		tnl = newTunnelClient(tunnel, "tcp")
		break
	case "udp-client":
		tnl = newTunnelClient(tunnel, "udp")
		break
	case "tcp-server":
		tnl = newTunnelTcpServer(tunnel)
		break
	case "udp-server":
		tnl = newTunnelUdpServer(tunnel)
		break
	default:
		return nil, fmt.Errorf("Unsupport type %s ", tunnel.Type)
	}
	return tnl, nil
}

func resolvePort(addr string) string {
	if strings.IndexByte(addr, ':') == -1 {
		return ":" + addr
	}
	return addr
}
