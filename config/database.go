package config

//Database 参数
type Database struct {
	Path string `json:"path"`
}

var DatabaseDefault = Database{
	Path: "iot-master-gateway.db",
}
