package main

import (
	"github.com/kardianos/service"
	"github.com/zgwit/iot-master/v2/pkg/log"
	"iot-master-gateway/args"
	"iot-master-gateway/config"
	"iot-master-gateway/db"
	"iot-master-gateway/dbus"
	"os"
	"os/signal"
	"syscall"
)

var serviceConfig = &service.Config{
	Name:        "iot-master-gateway",
	DisplayName: "物联大师网关",
	Description: "物联大师网关",
	Arguments:   nil,
}

func main() {
	args.Parse()

	//传递参数到服务
	serviceConfig.Arguments = []string{"-c", args.ConfigPath}

	// 构建服务对象
	program := &Program{}
	s, err := service.New(program, serviceConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 用于记录系统日志
	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if args.Uninstall {
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("卸载服务成功")
		return
	}

	if args.Install {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("安装服务成功")
		return
	}

	err = s.Run()
	if err != nil {
		_ = logger.Error(err)
	}
}

type Program struct{}

func (p *Program) Start(s service.Service) error {
	//log.Println("===开始服务===")
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	//log.Println("===停止服务===")
	_ = shutdown()
	return nil
}

func (p *Program) run() {

	// 此处编写具体的服务代码
	hup := make(chan os.Signal, 2)
	signal.Notify(hup, syscall.SIGHUP)
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, os.Kill)

	//原本的Main函数
	originMain()

	select {
	case <-hup:
	case <-quit:
		//优雅地结束
		_ = shutdown()
		//os.Exit(0)
	}
}

func originMain() {
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = log.Open(config.Config.Log)
	if err != nil {
		log.Fatal(err)
	}

	//加载数据库
	err = db.Open(config.Config.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbus.Open(config.Config.MQTT)
}

func shutdown() error {

	//_ = database.Close()
	//_ = tsdb.Close()
	//core.Close()
	//master.Close()

	//只关闭Web就行了，其他通过defer关闭

	return nil
}
