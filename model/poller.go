package model

// Poller 采集器
type Poller struct {
	Area    string `json:"area"`
	Address string `json:"address"`
	Length  int    `json:"length"`
}
