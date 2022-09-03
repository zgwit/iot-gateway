package model

import (
	"regexp"
	"time"
)

type Server struct {
	Id        uint64          `json:"id"`
	Name      string          `json:"name"`
	Type      string          `json:"type"` //tcp udp
	Addr      string          `json:"addr"`
	Register  RegisterPacket  `json:"register"`
	Heartbeat HeartBeatPacket `json:"heartbeat"`
	Protocol  Protocol        `json:"protocol"`
	Devices   []DefaultDevice `json:"devices"` //默认设备
	Disabled  bool            `json:"disabled"`
	Created   time.Time       `json:"created" xorm:"created"`
}

//RegisterPacket 注册包
type RegisterPacket struct {
	Regex  string `json:"regex,omitempty"`
	Length int    `json:"length,omitempty"`

	regex *regexp.Regexp
}
