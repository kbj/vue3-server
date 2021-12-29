package main

import (
	"github.com/gofiber/fiber/v2"
	"vue3-server/common"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		Prefork:      false,
		AppName:      "vue3-cms-server",
	})

	// 初始化
	common.Init(app)

	// 启动服务
	common.Start(app, ":8080")
}
