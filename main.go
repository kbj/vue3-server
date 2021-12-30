package main

import (
	"github.com/gofiber/fiber/v2"
	"vue3-server/common/boot"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		Prefork:      false,
		AppName:      "vue3-cms-server",
		ErrorHandler: boot.ErrorHandler(),
	})

	// 初始化
	boot.Init(app)

	// 启动服务
	boot.Start(app, ":8080")
}
