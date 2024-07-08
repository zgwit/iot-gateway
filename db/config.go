package db

import (
	"github.com/god-jason/bucket/config"
)

const MODULE = "database"

func init() {
	config.Register(MODULE, "type", "mysql")
	config.Register(MODULE, "url", ".db")
	config.Register(MODULE, "debug", false)
	config.Register(MODULE, "sync", true)
}
