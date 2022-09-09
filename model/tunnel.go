package model

import (
	"time"
)

type Protocol struct {
	Name    string                 `json:"name"`
	Options map[string]interface{} `json:"options"`
}

// Tunnel 通道模型
type Tunnel struct {
	Id       uint64        `json:"id"`
	ServerId uint64        `json:"server_id" boltholdIndex:"ServerId"`
	Name     string        `json:"name"`
	SN       string        `json:"sn" boltholdIndex:"SN"`
	Type     string        `json:"type"` //serial tcp-client tcp-server udp-client udp-server server-tcp server-udp
	Addr     string        `json:"addr"`
	Remote   string        `json:"remote"`
	Retry    Retry         `json:"retry"` //重试
	Serial   SerialOptions `json:"serial"`
	Protocol Protocol      `json:"protocol"`
	Last     time.Time     `json:"last"`
	Created  time.Time     `json:"created"`
}

type DefaultDevice struct {
	Station   int    `json:"station"`
	ProductId string `json:"product_id"`
}

type TunnelEx struct {
	Tunnel
	Running bool   `json:"running"`
	Online  bool   `json:"online"`
	Server  string `json:"server"`
}

func (tunnel *TunnelEx) TableName() string {
	return "tunnel"
}

type Retry struct {
	Enable  bool `json:"enable"`
	Timeout int  `json:"timeout"`
	Maximum int  `json:"maximum"`
}

// SerialOptions 串口参数
type SerialOptions struct {
	Port     string `json:"port"`      // /dev/tty.usb.. COM1
	BaudRate uint   `json:"baud_rate"` //9600 ... 115200 ...
	DataBits uint   `json:"data_bits"` //5 6 7 8
	StopBits uint   `json:"stop_bits"` //1 2
	Parity   uint   `json:"parity"`    // 0:NONE 1:ODD 2:EVEN
	//RS485    bool   `json:"rs485"`
}
