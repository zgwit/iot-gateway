package model

import "encoding/gob"

func init() {
	gob.Register(Tunnel{})
	gob.Register(Server{})
	gob.Register(Product{})
	gob.Register(Device{})
}
