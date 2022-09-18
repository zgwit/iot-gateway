package db

// Options 参数
type Options struct {
	Path string `json:"path"`
}

var Default = Options{
	Path: "iot-master-gateway.db",
}
