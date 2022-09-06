package core

import (
	"fmt"
	"iot-master-gateway/model"
)

// Server 通道
type Server interface {
	Open() error
	Close() error
	GetTunnel(id uint64) Tunnel
	Running() bool
}

// NewServer 创建通道
func NewServer(server *model.Server) (Server, error) {
	var svr Server
	switch server.Type {
	case "tcp":
		svr = newServerTCP(server)
		break
	case "udp":
		svr = newServerUDP(server)
		break
	default:
		return nil, fmt.Errorf("Unsupport type %s ", server.Type)
	}

	return svr, nil
}
