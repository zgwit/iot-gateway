package internal

import (
	"net"
	"time"
)

type netConn struct {
	net.Conn
}

func (c *netConn) SetReadTimeout(t time.Duration) error {
	return c.SetReadDeadline(time.Now().Add(t))
}
