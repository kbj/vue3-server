package main

import (
	"github.com/gin-gonic/gin"
	"vue3-server/common"
)

func main() {
	engine := gin.Default()

	// 初始化
	common.Init(engine)

	// 优雅关机
	common.GracefulShutDown(engine, ":8080")
}
