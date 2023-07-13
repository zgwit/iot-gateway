package types

import (
	"time"
)

type Tunnel struct {
	Id   string `json:"id,omitempty" xorm:"pk"` //ID
	Name string `json:"name,omitempty"`         //名称
	Desc string `json:"desc,omitempty"`         //说明

	Heartbeat string `json:"heartbeat,omitempty"` //心跳包

	Protocol `xorm:"extends"`
	Poller   `xorm:"extends"`

	Running bool `json:"running,omitempty" xorm:"-"`

	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" xorm:"created"` //创建时间
}

type Poller struct {
	PollerPeriod   uint `json:"poller_period,omitempty"`   //采集周期
	PollerInterval uint `json:"poller_interval,omitempty"` //采集间隔
}

type Protocol struct {
	ProtocolName    string `json:"protocol_name,omitempty"`    //协议 rtu tcp parallel-tcp
	ProtocolOptions string `json:"protocol_options,omitempty"` //协议参数
}

type Retry struct {
	RetryTimeout uint `json:"retry_timeout,omitempty"` //重试时间
	RetryMaximum uint `json:"retry_maximum,omitempty"` //最大次数
}
