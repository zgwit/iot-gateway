package db

import (
	"github.com/god-jason/bucket/config"
	"github.com/god-jason/bucket/lib"
)

const MODULE = "database"

func init() {
	config.Register(MODULE, "type", "sqlite")
	config.Register(MODULE, "url", lib.AppName()+".db")
	config.Register(MODULE, "debug", false)
	config.Register(MODULE, "sync", true)
}
