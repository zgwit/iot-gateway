package curd

import (
	"github.com/gin-gonic/gin"
	"github.com/zgwit/iot-gateway/db"
)

func ApiClear[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data T
		_, err := db.Engine.Where("1=1").Delete(&data)
		if err != nil {
			Error(ctx, err)
			return
		}
		OK(ctx, nil)
	}
}
