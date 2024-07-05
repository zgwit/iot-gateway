package main

import (
	"embed"
	_ "git.zgwit.com/gateway/dlt645"
	"github.com/god-jason/bucket/boot"
	"github.com/god-jason/bucket/pkg/service"
	"github.com/god-jason/bucket/web"
	"github.com/zgwit/iot-gateway/api"
	"github.com/zgwit/iot-gateway/args"
	_ "github.com/zgwit/iot-gateway/client"
	_ "github.com/zgwit/iot-gateway/device"
	_ "github.com/zgwit/iot-gateway/modbus"
	_ "github.com/zgwit/iot-gateway/product"
	_ "github.com/zgwit/iot-gateway/serial"
	_ "github.com/zgwit/iot-gateway/server"
	"log"
	"net/http"
)

// go: embed all:www
var wwwFiles embed.FS

func main() {
	args.Parse()

	err := service.Register(Startup, Shutdown)
	if err != nil {
		log.Fatal(err)
	}

	if args.Uninstall {
		err = service.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("卸载服务成功")
		return
	}

	if args.Install {
		err = service.Install()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("安装服务成功")
		return
	}

	err = service.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Startup() error {
	err := boot.Startup()
	if err != nil {
		return err
	}

	//注册前端接口
	api.RegisterRoutes(web.Engine.Group("/api"))

	//阻塞
	return web.Serve()
}

func Static(fs *web.FileSystem) {
	//前端静态文件
	fs.Put("", http.FS(wwwFiles), "www", "www/index.html")
}

func Shutdown() error {
	return boot.Shutdown()
}
