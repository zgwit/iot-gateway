package connect

import (
	"fmt"
	"io"
	"iot-master-gateway/model"
	"iot-master-gateway/pkg/events"
	"strings"
	"time"
)

//Tunnel 通道
type Tunnel interface {
	events.EventInterface

	Model() *model.Tunnel

	Write(data []byte) error

	Open() error

	Close() error

	Running() bool

	Online() bool

	First() bool

	//Pipe 透传
	Pipe(pipe io.ReadWriteCloser)

	//Ask 发送指令，接收数据
	Ask(cmd []byte, timeout time.Duration) ([]byte, error)
}

//NewTunnel 创建通道
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
