package db

import (
	"iot-master-gateway/config"
	"iot-master-gateway/model"
	"xorm.io/xorm"
	"xorm.io/xorm/log"

	//加载数据库驱动
	//_ "github.com/mattn/go-sqlite3" //CGO版本
	_ "github.com/glebarez/go-sqlite" //纯Go版本 使用ccgo翻译的，暂未发现问题
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var Engine *xorm.Engine

func Open(cfg *config.Database) error {
	var err error
	Engine, err = xorm.NewEngine(cfg.Type, cfg.URL)
	if err != nil {
		return err
	}
	if cfg.Debug {
		Engine.ShowSQL(true)
	}

	Engine.SetLogLevel(log.LogLevel(cfg.LogLevel))
	//Engine.SetLogger(logrus.StandardLogger())

	//同步表
	if cfg.Sync {
		err = Sync()
		if err != nil {
			return err
		}
	}

	return nil
}

func Close() error {
	return Engine.Close()
}

func Sync() error {
	return Engine.Sync2(
		new(model.User), new(model.Password),
		new(model.Tunnel), new(model.Server),
		new(model.Device), new(model.Product),
		new(model.Hmi), new(model.Component),
		new(model.Event),
	)
}
