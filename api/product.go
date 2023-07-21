package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-gateway/internal"
	"github.com/zgwit/iot-gateway/types"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询产品数量
// @Schemes
// @Description 查询产品数量
// @Tags product
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回产品数量
// @Router /product/count [post]
func noopProductCount() {}

// @Summary 查询产品
// @Schemes
// @Description 查询产品
// @Tags product
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Product] 返回产品信息
// @Router /product/search [post]
func noopProductSearch() {}

// @Summary 查询产品
// @Schemes
// @Description 查询产品
// @Tags product
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Product] 返回产品信息
// @Router /product/list [get]
func noopProductList() {}

// @Summary 创建产品
// @Schemes
// @Description 创建产品
// @Tags product
// @Param search body types.Product true "产品信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Product] 返回产品信息
// @Router /product/create [post]
func noopProductCreate() {}

// @Summary 修改产品
// @Schemes
// @Description 修改产品
// @Tags product
// @Param id path int true "产品ID"
// @Param product body types.Product true "产品信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Product] 返回产品信息
// @Router /product/{id} [post]
func noopProductUpdate() {}

// @Summary 获取产品
// @Schemes
// @Description 获取产品
// @Tags product
// @Param id path int true "产品ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Product] 返回产品信息
// @Router /product/{id} [get]
func noopProductGet() {}

// @Summary 删除产品
// @Schemes
// @Description 删除产品
// @Tags product
// @Param id path int true "产品ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Product] 返回产品信息
// @Router /product/{id}/delete [get]
func noopProductDelete() {}

// @Summary 导出产品
// @Schemes
// @Description 导出产品
// @Tags product
// @Accept json
// @Produce octet-stream
// @Router /product/export [get]
func noopProductExport() {}

// @Summary 导入产品
// @Schemes
// @Description 导入产品
// @Tags product
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回产品数量
// @Router /product/import [post]
func noopProductImport() {}

func productRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Product]())

	app.POST("/search", curd.ApiSearch[types.Product]("id", "name", "desc", "created"))

	app.GET("/list", curd.ApiList[types.Product]("id", "name", "desc", "created"))

	app.POST("/create", curd.ApiCreateHook[types.Product](curd.GenerateRandomId[types.Product](8), func(m *types.Product) error {
		return internal.LoadProduct(m)
	}))

	app.GET("/:id", curd.ParseParamStringId, curd.ApiGet[types.Product]())

	app.POST("/:id", curd.ParseParamStringId, curd.ApiUpdateHook[types.Product](nil, func(m *types.Product) error {
		return internal.LoadProduct(m)
	}, "id", "name", "desc", "mappers", "filters", "calculators"))

	app.GET("/:id/delete", curd.ParseParamStringId, curd.ApiDeleteHook[types.Product](nil, func(id interface{}) error {
		//internal.DeleteProduct(id)
		return nil
	}))

	app.GET("/export", curd.ApiExport("product", "product"))

	app.POST("/import", curd.ApiImport("product"))
}
