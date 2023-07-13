package connect

import (
	"github.com/iot-master-contrib/modbus/types"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/lib"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"xorm.io/xorm"
)

var serials lib.Map[Serial]
var clients lib.Map[Client]
var servers lib.Map[Server]
var links lib.Map[Link]

func LoadSerials() error {
	var serials []*types.Serial
	err := db.Engine.Find(&serials)
	if err != nil {
		if err == xorm.ErrNotExist {
			return nil
		}
		return err
	}
	for _, m := range serials {
		if m.Disabled {
			continue
		}
		go func(m *types.Serial) {
			err := LoadSerial(m)
			if err != nil {
				log.Error(err)
			}
		}(m)
	}
	return nil
}

func LoadSerial(m *types.Serial) error {
	s := NewSerial(m)
	serials.Store(m.Id, s)
	return s.Open()
}

func GetSerial(id string) *Serial {
	return serials.Load(id)
}

func LoadClients() error {
	var clients []*types.Client
	err := db.Engine.Find(&clients)
	if err != nil {
		if err == xorm.ErrNotExist {
			return nil
		}
		return err
	}
	for _, m := range clients {
		if m.Disabled {
			continue
		}
		go func(m *types.Client) {
			err := LoadClient(m)
			if err != nil {
				log.Error(err)
			}
		}(m)
	}
	return nil
}

func LoadClient(m *types.Client) error {
	s := NewClient(m)
	clients.Store(m.Id, s)
	return s.Open()
}

func GetClient(id string) *Client {
	return clients.Load(id)
}

func LoadServers() error {
	var servers []*types.Server
	err := db.Engine.Find(&servers)
	if err != nil {
		if err == xorm.ErrNotExist {
			return nil
		}
		return err
	}
	for _, m := range servers {
		if m.Disabled {
			continue
		}
		go func(m *types.Server) {
			err := LoadServer(m)
			if err != nil {
				log.Error(err)
			}
		}(m)
	}
	return nil
}

func LoadServer(m *types.Server) error {
	s := NewServer(m)
	servers.Store(m.Id, s)
	return s.Open()
}

func GetServer(id string) *Server {
	return servers.Load(id)
}

func GetLink(id string) *Link {
	return links.Load(id)
}

func Load() error {
	err := LoadSerials()
	if err != nil {
		return err
	}
	err = LoadClients()
	if err != nil {
		return err
	}
	err = LoadServers()
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	links.Range(func(name string, link *Link) bool {
		err := link.Close()
		if err != nil {
			log.Error(err)
		}
		return true
	})
}
