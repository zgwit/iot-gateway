package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-gateway/internal"
	"github.com/zgwit/iot-gateway/types"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
)

// @Summary 查询服务器数量
// @Schemes
// @Description 查询服务器数量
// @Tags server
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回服务器数量
// @Router /server/count [post]
func noopServerCount() {}

// @Summary 查询服务器
// @Schemes
// @Description 查询服务器
// @Tags server
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Server] 返回服务器信息
// @Router /server/search [post]
func noopServerSearch() {}

// @Summary 查询服务器
// @Schemes
// @Description 查询服务器
// @Tags server
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Server] 返回服务器信息
// @Router /server/list [get]
func noopServerList() {}

// @Summary 创建服务器
// @Schemes
// @Description 创建服务器
// @Tags server
// @Param search body types.Server true "服务器信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Server] 返回服务器信息
// @Router /server/create [post]
func noopServerCreate() {}

// @Summary 修改服务器
// @Schemes
// @Description 修改服务器
// @Tags server
// @Param id path int true "服务器ID"
// @Param server body types.Server true "服务器信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Server] 返回服务器信息
// @Router /server/{id} [post]
func noopServerUpdate() {}

// @Summary 获取服务器
// @Schemes
// @Description 获取服务器
// @Tags server
// @Param id path int true "服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Server] 返回服务器信息
// @Router /server/{id} [get]
func noopServerGet() {}

// @Summary 删除服务器
// @Schemes
// @Description 删除服务器
// @Tags server
// @Param id path int true "服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Server] 返回服务器信息
// @Router /server/{id}/delete [get]
func noopServerDelete() {}

// @Summary 启用服务器
// @Schemes
// @Description 启用服务器
// @Tags server
// @Param id path int true "服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Server] 返回服务器信息
// @Router /server/{id}/enable [get]
func noopServerEnable() {}

// @Summary 禁用服务器
// @Schemes
// @Description 禁用服务器
// @Tags server
// @Param id path int true "服务器ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Server] 返回服务器信息
// @Router /server/{id}/disable [get]
func noopServerDisable() {}

// @Summary 启动服务端
// @Schemes
// @Description 启动服务端
// @Tags server
// @Param id path int true "服务端ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Server] 返回服务端信息
// @Router /server/{id}/start [get]
func noopServerStart() {}

// @Summary 停止服务端
// @Schemes
// @Description 停止服务端
// @Tags server
// @Param id path int true "服务端ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Server] 返回服务端信息
// @Router /server/{id}/stop [get]
func noopServerStop() {}

// @Summary 导出服务端
// @Schemes
// @Description 导出服务端
// @Tags product
// @Accept json
// @Produce octet-stream
// @Router /server/export [get]
func noopServerExport() {}

// @Summary 导入服务端
// @Schemes
// @Description 导入服务端
// @Tags product
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回服务端数量
// @Router /server/import [post]
func noopServerImport() {}

func serverRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Server]())

	app.POST("/search", curd.ApiSearchHook[types.Server](func(servers []*types.Server) error {
		for k, server := range servers {
			c := internal.GetServer(server.Id)
			if c != nil {
				servers[k].Running = c.Running()
			}
		}
		return nil
	}))

	app.GET("/list", curd.ApiList[types.Server]())

	app.POST("/create", curd.ApiCreateHook[types.Server](curd.GenerateRandomId[types.Server](8), func(value *types.Server) error {
		return internal.LoadServer(value)
	}))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGetHook[types.Server](func(server *types.Server) error {
		c := internal.GetServer(server.Id)
		if c != nil {
			server.Running = c.Running()
		}
		return nil
	}))

	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[types.Server](nil, func(value *types.Server) error {
		c := internal.GetServer(value.Id)
		err := c.Close()
		if err != nil {
			log.Error(err)
		}
		return internal.LoadServer(value)
	},
		"id", "name", "desc", "heartbeat", "poller_period", "poller_interval", "protocol_name", "protocol_options", "retry", "options", "disabled", "port", "standalone", "servers"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[types.Server](nil, func(value interface{}) error {
		id := value.(string)
		c := internal.GetServer(id)
		return c.Close()
	}))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisableHook[types.Server](true, nil, func(value interface{}) error {
		id := value.(string)
		c := internal.GetServer(id)
		return c.Close()
	}))

	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisableHook[types.Server](false, nil, func(value interface{}) error {
		id := value.(string)
		var m types.Server
		has, err := db.Engine.ID(id).Get(&m)
		if err != nil {
			return err
		}
		if !has {
			return fmt.Errorf("找不到 %s", id)
		}
		return internal.LoadServer(&m)
	}))

	app.GET(":id/start", curd.ParseParamStringId, func(ctx *gin.Context) {
		id := ctx.GetString("id")
		c := internal.GetServer(id)
		if c == nil {
			curd.Fail(ctx, "找不到连接")
			return
		}
		err := c.Open()
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	})

	app.GET(":id/stop", curd.ParseParamStringId, func(ctx *gin.Context) {
		id := ctx.GetString("id")
		c := internal.GetServer(id)
		if c == nil {
			curd.Fail(ctx, "找不到连接")
			return
		}
		err := c.Close()
		if err != nil {
			curd.Error(ctx, err)
			return
		}
		curd.OK(ctx, nil)
	})

	app.GET("/export", curd.ApiExport("server", "server"))
	app.POST("/import", curd.ApiImport("server"))

}
