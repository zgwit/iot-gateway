package core

import (
	"github.com/timshannon/bolthold"
	"github.com/zgwit/iot-master/v2/model"
	"github.com/zgwit/iot-master/v2/pkg/log"
	"iot-master-gateway/connect"
	"iot-master-gateway/db"
	"iot-master-gateway/dbus"
)

var Products Map[Product]
var Devices Map[Device]
var Tunnels Map[Tunnel]
var Servers Map[Server]

func Open(node string) {
	RegisterGatewayHandler(dbus.MQTT, node)

	err := loadProducts()
	if err != nil {
		log.Error(err)
	}

	err = loadTunnels()
	if err != nil {
		log.Error(err)
	}

	err = loadServers()
	if err != nil {
		log.Error(err)
	}
}

func loadTunnels() error {
	return db.Store().ForEach(bolthold.Where("ServerId").Eq(""), func(tunnel *model.Tunnel) error {
		err := loadTunnel(tunnel)
		if err != nil {
			log.Error(err)
		}
		return nil
	})
}

func loadTunnel(tunnel *model.Tunnel) error {
	tnl, err := connect.NewTunnel(tunnel)
	if err != nil {
		return err
	}

	//加载设备
	devices, err := loadDevices(tunnel.Id)
	if err != nil {
		return err
	}

	Tunnels.Store(tunnel.Id, &Tunnel{
		tunnel:  tnl,
		adapter: newAdapter(tunnel.Protocol.Name, tunnel.Protocol.Options, tnl),
		devices: devices,
	})

	return nil
}

func loadServers() error {
	return db.Store().ForEach(nil, func(server *model.Server) error {
		err := loadServer(server)
		if err != nil {
			log.Error(err)
		}
		return nil
	})
}

func loadServer(server *model.Server) error {
	svr, err := connect.NewServer(server)
	if err != nil {
		return err
	}

	Servers.Store(server.Id, &Server{
		server: svr,
	})

	//TODO 监听新连接

	return nil
}

func loadDevices(tunnelId string) ([]*Device, error) {
	devices := make([]*Device, 0)
	err := db.Store().ForEach(bolthold.Where("TunnelId").Eq(tunnelId), func(device *model.Device) error {
		dev, err := loadDevice(device)
		if err != nil {
			log.Error(err)
		}
		devices = append(devices, dev)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return devices, nil
}

func loadDevice(device *model.Device) (*Device, error) {
	dev := &Device{
		Device:  device,
		product: getProduct(device.ProductId),
	}

	Devices.Store(device.Id, dev)

	return dev, nil
}

func loadProducts() error {
	return db.Store().ForEach(nil, func(product *model.Product) error {
		err := loadProduct(product)
		if err != nil {
			log.Error(err)
		}
		return nil
	})
}

func loadProduct(product *model.Product) error {
	pro := &Product{
		Product: product,
	}
	Products.Store(product.Id, pro)

	//TODO 找到协议，解析地址

	return nil
}

func getProduct(id string) *Product {
	pro := Products.Load(id)
	if pro != nil {
		return pro
	}

	var product model.Product
	err := db.Store().Get(id, &product)
	if err != nil {
		log.Error(err)
	}

	err = loadProduct(&product)
	if err != nil {
		log.Error(err)
		return nil
	}

	return Products.Load(id)
}
