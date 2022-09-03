package config

// MQTT 参数
type MQTT struct {
	Broker   string
	ClientID string
	Username string
	Password string
}

var MQTTDefault = MQTT{}
