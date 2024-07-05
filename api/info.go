package api

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/pkg/build"
	"github.com/zgwit/iot-gateway/curd"
	"runtime"
)

func info(ctx *gin.Context) {
	curd.OK(ctx, gin.H{
		"version": build.Version,
		"build":   build.Build,
		"git":     build.GitHash,
		"gin":     gin.Version,
		"runtime": runtime.Version(),
	})
}
