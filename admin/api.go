package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-gateway/api"
	"github.com/zgwit/iot-gateway/curd"
)

func init() {
	api.Register("GET", "me", me)
	api.Register("GET", "logout", logout)
	api.Register("POST", "password", password)
}

func me(ctx *gin.Context) {
	id := ctx.GetString("user")

	if id == "" {
		curd.Fail(ctx, "未登录")
		return
	}

	curd.OK(ctx, gin.H{"id": id})
}
