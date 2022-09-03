package model

import "encoding/gob"

func init() {
	gob.Register(User{})
	gob.Register(Product{})
	gob.Register(Device{})
	gob.Register(Product{})
	gob.Register(Tunnel{})
	gob.Register(Server{})
	gob.Register(Plugin{})
}
