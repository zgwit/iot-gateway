package core

import (
	"github.com/zgwit/go-plc/modbus"
	"github.com/zgwit/go-plc/protocol"
	"io"
)

var protocols = []protocol.Manifest{
	modbus.ManifestRTU,
	modbus.ManifestTCP,
}

func newAdapter(name string, opts string, conn io.ReadWriter) protocol.Protocol {
	for i := 0; i < len(protocols); i++ {
		if protocols[i].Name == name {
			return protocols[i].Factory(conn, opts)
		}
	}
	return nil
}
