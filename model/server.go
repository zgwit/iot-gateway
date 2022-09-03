package model

import (
	"regexp"
	"time"
)

type Server struct {
	Id        int64           `json:"id"`
	Name      string          `json:"name"`
	Type      string          `json:"type"` //tcp udp
	Addr      string          `json:"addr"`
	Register  RegisterPacket  `json:"register" xorm:"JSON"`
	Heartbeat HeartBeatPacket `json:"heartbeat" xorm:"JSON"`
	Protocol  Protocol        `json:"protocol" xorm:"JSON"`
	Devices   []DefaultDevice `json:"devices"` //默认设备
	Disabled  bool            `json:"disabled"`
	Updated   time.Time       `json:"updated" xorm:"updated"`
	Created   time.Time       `json:"created" xorm:"created"`
	//Deleted   time.Time       `json:"-" xorm:"deleted"`
}

type ServerEx struct {
	Server  `xorm:"extends"`
	Running bool `json:"running"`
}

func (s *ServerEx) TableName() string {
	return "server"
}

//RegisterPacket 注册包
type RegisterPacket struct {
	Regex  string `json:"regex,omitempty"`
	Length int    `json:"length,omitempty"`

	regex *regexp.Regexp
}
