package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iot-master-contrib/gateway/internal"
	"github.com/iot-master-contrib/gateway/types"
	"github.com/zgwit/iot-master/v3/pkg/curd"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
)

// @Summary 查询客户端数量
// @Schemes
// @Description 查询客户端数量
// @Tags client
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回客户端数量
// @Router /client/count [post]
func noopClientCount() {}

// @Summary 查询客户端
// @Schemes
// @Description 查询客户端
// @Tags client
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Client] 返回客户端信息
// @Router /client/search [post]
func noopClientSearch() {}

// @Summary 查询客户端
// @Schemes
// @Description 查询客户端
// @Tags client
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Client] 返回客户端信息
// @Router /client/list [get]
func noopClientList() {}

// @Summary 创建客户端
// @Schemes
// @Description 创建客户端
// @Tags client
// @Param search body types.Client true "客户端信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Client] 返回客户端信息
// @Router /client/create [post]
func noopClientCreate() {}

// @Summary 修改客户端
// @Schemes
// @Description 修改客户端
// @Tags client
// @Param id path int true "客户端ID"
// @Param client body types.Client true "客户端信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Client] 返回客户端信息
// @Router /client/{id} [post]
func noopClientUpdate() {}

// @Summary 获取客户端
// @Schemes
// @Description 获取客户端
// @Tags client
// @Param id path int true "客户端ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Client] 返回客户端信息
// @Router /client/{id} [get]
func noopClientGet() {}

// @Summary 删除客户端
// @Schemes
// @Description 删除客户端
// @Tags client
// @Param id path int true "客户端ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Client] 返回客户端信息
// @Router /client/{id}/delete [get]
func noopClientDelete() {}

// @Summary 启用客户端
// @Schemes
// @Description 启用客户端
// @Tags client
// @Param id path int true "客户端ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Client] 返回客户端信息
// @Router /client/{id}/enable [get]
func noopClientEnable() {}

// @Summary 禁用客户端
// @Schemes
// @Description 禁用客户端
// @Tags client
// @Param id path int true "客户端ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Client] 返回客户端信息
// @Router /client/{id}/disable [get]
func noopClientDisable() {}

// @Summary 启动客户端
// @Schemes
// @Description 启动客户端
// @Tags client
// @Param id path int true "客户端ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Client] 返回客户端信息
// @Router /client/{id}/start [get]
func noopClientStart() {}

// @Summary 停止客户端
// @Schemes
// @Description 停止客户端
// @Tags client
// @Param id path int true "客户端ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Client] 返回客户端信息
// @Router /client/{id}/stop [get]
func noopClientStop() {}

// @Summary 导出客户端
// @Schemes
// @Description 导出客户端
// @Tags product
// @Accept json
// @Produce octet-stream
// @Router /client/export [get]
func noopClientExport() {}

// @Summary 导入客户端
// @Schemes
// @Description 导入客户端
// @Tags product
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回客户端数量
// @Router /client/import [post]
func noopClientImport() {}

func clientRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Client]())

	app.POST("/search", curd.ApiSearchHook[types.Client](func(clients []*types.Client) error {
		for k, client := range clients {
			c := internal.GetClient(client.Id)
			if c != nil {
				clients[k].Running = c.Running()
			}
		}
		return nil
	}))

	app.GET("/list", curd.ApiList[types.Client]())

	app.POST("/create", curd.ApiCreateHook[types.Client](curd.GenerateRandomId[types.Client](8), func(value *types.Client) error {
		return internal.LoadClient(value)
	}))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGetHook[types.Client](func(client *types.Client) error {
		c := internal.GetClient(client.Id)
		if c != nil {
			client.Running = c.Running()
		}
		return nil
	}))

	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[types.Client](nil, func(value *types.Client) error {
		c := internal.GetClient(value.Id)
		err := c.Close()
		if err != nil {
			log.Error(err)
		}
		return internal.LoadClient(value)
	},
		"id", "name", "desc", "heartbeat", "poller_period", "poller_interval", "protocol_name", "protocol_options", "disabled", "retry_timeout", "retry_maximum", "net", "addr", "port"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[types.Client](nil, func(value interface{}) error {
		id := value.(string)
		c := internal.GetClient(id)
		return c.Close()
	}))

	app.GET(":id/disable", curd.ParseParamStringId, curd.ApiDisableHook[types.Client](true, nil, func(value interface{}) error {
		id := value.(string)
		c := internal.GetClient(id)
		return c.Close()
	}))

	app.GET(":id/enable", curd.ParseParamStringId, curd.ApiDisableHook[types.Client](false, nil, func(value interface{}) error {
		id := value.(string)
		var m types.Client
		has, err := db.Engine.ID(id).Get(&m)
		if err != nil {
			return err
		}
		if !has {
			return fmt.Errorf("找不到 %s", id)
		}
		return internal.LoadClient(&m)
	}))

	app.GET(":id/start", curd.ParseParamStringId, func(ctx *gin.Context) {
		id := ctx.GetString("id")
		c := internal.GetClient(id)
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
		c := internal.GetClient(id)
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

	app.GET("/export", curd.ApiExport("client", "client"))

	app.POST("/import", curd.ApiImport("client"))

}
