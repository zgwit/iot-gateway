package protocols

import (
	"fmt"
	"github.com/iot-master-contrib/gateway/connect"
	"github.com/iot-master-contrib/gateway/protocols/define"
	"github.com/iot-master-contrib/gateway/types"
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
