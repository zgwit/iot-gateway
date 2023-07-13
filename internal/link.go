package internal

import (
	"errors"
	"github.com/iot-master-contrib/gateway/types"
	"net"
)

// Link 网络连接
type Link struct {
	tunnelBase
	model *types.Link
}

func newLink(client *types.Link, conn net.Conn) *Link {
	return &Link{
		model: client,
		tunnelBase: tunnelBase{
			Conn: &netConn{conn},
		}}
}

func (l *Link) Open() error {
	return errors.New("conn cannot open")
}
