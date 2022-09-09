package model

import (
	"time"
)

type Server struct {
	Id       uint64          `json:"id"`
	Name     string          `json:"name"`
	Type     string          `json:"type"` //tcp udp
	Addr     string          `json:"addr"`
	Protocol Protocol        `json:"protocol"`
	Devices  []DefaultDevice `json:"devices"` //默认设备
	Disabled bool            `json:"disabled"`
	Created  time.Time       `json:"created"`
}
