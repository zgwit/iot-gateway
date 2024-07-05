package modbus

import (
	"github.com/zgwit/iot-gateway/db"
)

func init() {
	db.Register(new(Device))
}

type Station struct {
	Slave uint8 `json:"slave"`
}

type Device struct {
	Id        string `json:"id" xorm:"pk"`
	ProductId string `json:"product_id"`

	//modbus站号
	//ModbusStation uint8 `json:"modbus_station,omitempty"`
	Station Station `json:"station,omitempty" xorm:"json"`

	//映射和轮询表
	pollers *[]*Poller

	mapper *Mapper
}
