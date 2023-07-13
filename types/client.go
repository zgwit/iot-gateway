package types

type Client struct {
	Tunnel       `xorm:"extends"`
	RetryOptions `xorm:"extends"`
	Net          string `json:"net,omitempty"`  //类型 tcp udp
	Addr         string `json:"addr,omitempty"` //地址，主机名或IP
	Port         uint16 `json:"port,omitempty"` //端口号
}
