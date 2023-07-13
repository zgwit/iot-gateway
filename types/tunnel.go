package types

import (
	"time"
)

type Tunnel struct {
	Id   string `json:"id,omitempty" xorm:"pk"` //ID
	Name string `json:"name,omitempty"`         //名称
	Desc string `json:"desc,omitempty"`         //说明

	Heartbeat string `json:"heartbeat,omitempty"` //心跳包

	ProtocolOptions `xorm:"extends"`
	PollerOptions   `xorm:"extends"`

	Running bool `json:"running,omitempty" xorm:"-"`

	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" xorm:"created"` //创建时间
}

type PollerOptions struct {
	Period   uint `json:"poller_period,omitempty"`   //采集周期
	Interval uint `json:"poller_interval,omitempty"` //采集间隔
}

type ProtocolOptions struct {
	Name    string         `json:"protocol_name,omitempty"`    //协议 rtu tcp parallel-tcp
	Options map[string]any `json:"protocol_options,omitempty"` //协议参数
}

type RetryOptions struct {
	Timeout uint `json:"retry_timeout,omitempty"` //重试时间
	Maximum uint `json:"retry_maximum,omitempty"` //最大次数
}
