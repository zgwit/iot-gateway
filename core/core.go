package core

import "iot-master-gateway/dbus"

var Products Map[Product]
var Devices Map[Device]
var Tunnels Map[Tunnel]
var Servers Map[Server]

func Open(node string) {
	RegisterGatewayHandler(dbus.MQTT, node)

}
