package modbus

import (
	"github.com/iot-master-contrib/gateway/connect"
	"github.com/iot-master-contrib/gateway/protocols"
	"github.com/iot-master-contrib/gateway/protocols/define"
	"github.com/iot-master-contrib/gateway/types"
)

var code = []types.Code{
	{Code: 1, Label: "线圈"},
	{Code: 2, Label: "离散输入"},
	{Code: 3, Label: "保持寄存器"},
	{Code: 4, Label: "输入寄存器"},
}

var modbusRtu = define.Protocol{
	Name:    "modbus-rtu",
	Label:   "Modbus RTU",
	Codes:   code,
	Factory: createRTU,
}

var modbusTCP = define.Protocol{
	Name:    "modbus-tcp",
	Label:   "Modbus TCP",
	Codes:   code,
	Factory: createRTU,
}

func createRTU(conn connect.Conn, opts map[string]any) define.Adapter {
	return define.Adapter(NewRTU(conn, opts))
}

func init() {
	protocols.Register(modbusRtu)
	//protocols.Register(modbusTCP)
}
