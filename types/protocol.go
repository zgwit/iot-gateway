package types

type Station map[string]int

type Address struct {
	Code int `json:"code,omitempty"`
	Addr int `json:"addr,omitempty"`
	Area int `json:"area,omitempty"`
}

type Code struct {
	Code  int    `json:"code,omitempty"`
	Label string `json:"label,omitempty"`
	IsBit bool   `json:"isBit,omitempty"`
}
