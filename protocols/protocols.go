package protocols

import (
	"fmt"
	"github.com/zgwit/iot-gateway/connect"
	"github.com/zgwit/iot-gateway/protocols/define"
	"github.com/zgwit/iot-gateway/types"
)

var protocols map[string]define.Protocol

func Protocols() []define.Protocol {
	var ps []define.Protocol
	for _, p := range protocols {
		ps = append(ps, p)
	}
	return ps
}

func Register(protocol define.Protocol) {
	protocols[protocol.Name] = protocol
}

func Create(conn connect.Conn, name string, opts types.Options) (define.Adapter, error) {
	if p, ok := protocols[name]; ok {
		return p.Factory(conn, opts), nil
	}
	return nil, fmt.Errorf("协议 %s 找不到", name)
}
