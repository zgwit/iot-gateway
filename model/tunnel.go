package model

import (
	"bytes"
	"encoding/hex"
	"regexp"
	"time"
)

type Protocol struct {
	Name    string                 `json:"name"`
	Options map[string]interface{} `json:"options"`
}

//Tunnel 通道模型
type Tunnel struct {
	Id        int64           `json:"id"`
	ServerId  int64           `json:"server_id" xorm:"index"`
	Name      string          `json:"name"`
	Type      string          `json:"type"` //serial tcp-client tcp-server udp-client udp-server server-tcp server-udp
	Addr      string          `json:"addr" xorm:"index"`
	Remote    string          `json:"remote"`
	Retry     Retry           `json:"retry" xorm:"JSON"` //重试
	Heartbeat HeartBeatPacket `json:"heartbeat" xorm:"JSON"`
	Serial    SerialOptions   `json:"serial" xorm:"JSON"`
	Protocol  Protocol        `json:"protocol" xorm:"JSON"`
	//Devices   []DefaultDevice  `json:"devices"` //默认设备
	Disabled bool      `json:"disabled"`
	Last     time.Time `json:"last"`
	Updated  time.Time `json:"updated" xorm:"updated"`
	Created  time.Time `json:"created" xorm:"created"`
	//Deleted  time.Time `json:"-" xorm:"deleted"`
}

type DefaultDevice struct {
	Station   int    `json:"station"`
	ProductId string `json:"product_id"`
}

type TunnelEx struct {
	Tunnel  `xorm:"extends"`
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

//SerialOptions 串口参数
type SerialOptions struct {
	Port     string `json:"port"`      // /dev/tty.usb.. COM1
	BaudRate uint   `json:"baud_rate"` //9600 ... 115200 ...
	DataBits uint   `json:"data_bits"` //5 6 7 8
	StopBits uint   `json:"stop_bits"` //1 2
	Parity   uint   `json:"parity"`    // 0:NONE 1:ODD 2:EVEN
	//RS485    bool   `json:"rs485"`
}

//Check 检查
func (p *RegisterPacket) Check(buf []byte) bool {
	if p.Regex != "" {
		if p.regex == nil {
			p.regex = regexp.MustCompile(p.Regex)
		}
		return p.regex.MatchString(string(buf))
	}
	if p.Length > 0 {
		if len(buf) != p.Length {
			return false
		}
	}
	return true
}

//HeartBeatPacket 心跳包
type HeartBeatPacket struct {
	Enable  bool   `json:"enable"`
	Timeout int64  `json:"timeout"`
	Regex   string `json:"regex,omitempty"`
	Hex     string `json:"hex,omitempty"`
	Text    string `json:"text,omitempty"`
	//Length  int    `json:"length,omitempty"`

	hex   []byte
	regex *regexp.Regexp
	last  int64
}

//Check 检查
func (p *HeartBeatPacket) Check(buf []byte) bool {

	now := time.Now().Unix()
	if p.last == 0 {
		p.last = now
	}
	if p.last+p.Timeout > now {
		p.last = now
		return false
	}
	p.last = now

	if p.Regex != "" {
		if p.regex == nil {
			p.regex = regexp.MustCompile(p.Regex)
		}
		return p.regex.Match(buf)
	}

	//if p.Length > 0 {
	//	if len(buf) != p.Length {
	//		return false
	//	}
	//}

	if p.Hex != "" {
		if p.hex == nil {
			//var err error
			p.hex, _ = hex.DecodeString(p.Hex)
		}
		return bytes.Equal(p.hex, buf)
	}

	if p.Text != "" {
		return p.Text == string(buf)
	}

	return true
}
