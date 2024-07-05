package serial

import (
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"github.com/zgwit/iot-gateway/db"
	"xorm.io/xorm"
)

var serials lib.Map[Serial]

func LoadSerials() error {
	var serials []*Serial
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
		go func(m *Serial) {
			err := LoadSerial(m)
			if err != nil {
				log.Error(err)
			}
		}(m)
	}
	return nil
}

func LoadSerial(m *Serial) error {
	serials.Store(m.Id, m)
	return m.Open()
}

func GetSerial(id string) *Serial {
	return serials.Load(id)
}
