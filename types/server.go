package types

type Server struct {
	Tunnel     `xorm:"extends"`
	Port       uint16          `json:"port,omitempty"`       //监听端口
	Standalone bool            `json:"standalone,omitempty"` //单例模式（不支持注册）
	Devices    []DefaultDevice `json:"devices,omitempty" xorm:"json"`
}

type DefaultDevice struct {
	Slave     uint8  `json:"slave"`          //从站号
	Name      string `json:"name,omitempty"` //名称
	ProductId string `json:"product_id"`     //产品ID
}

type Link struct {
	Tunnel   `xorm:"extends"`
	ServerId string `json:"server_id" xorm:"index"` //服务器ID
	Remote   string `json:"remote,omitempty"`       //远程地址
}
