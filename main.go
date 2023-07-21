package gateway

import (
	"embed"
	"github.com/zgwit/iot-gateway/api"
	_ "github.com/zgwit/iot-gateway/docs"
	"github.com/zgwit/iot-gateway/internal"
	"github.com/zgwit/iot-gateway/types"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"net/http"
)

//go:embed all:www
var wwwFiles embed.FS

// @title 物联大师网关接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /api/gateway/api/
// @query.collection.format multi
func main() {
}

func Startup(app *web.Engine) error {

	//同步表结构
	err := db.Engine.Sync2(
		new(types.Client), new(types.Server),
		new(types.Link), new(types.Serial),
		new(types.Product), new(types.Device),
	)
	if err != nil {
		log.Fatal(err)
	}

	//内部加载
	err = internal.LoadProducts()
	if err != nil {
		log.Fatal(err)
	}

	//连接
	err = internal.Load()
	if err != nil {
		log.Fatal(err)
	}
	//defer connect.Close()

	//注册前端接口
	api.RegisterRoutes(app.Group("/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(&app.RouterGroup, "gateway")

	return nil
}

func Static(fs *web.FileSystem) {
	//前端静态文件
	fs.Put("", http.FS(wwwFiles), "www", "www/index.html")
}

func Shutdown() error {

	//只关闭Web就行了，其他通过defer关闭

	return nil
}
