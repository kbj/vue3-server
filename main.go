package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"vue3-server/common/boot"
	"vue3-server/common/global"
)

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		Prefork:      false,
		ErrorHandler: boot.ErrorHandler(),
	})

	// 初始化
	boot.Init(app)

	// 启动服务
	boot.Start(app, fmt.Sprintf("%s:%s", global.Config.System.Listen, global.Config.System.Port))
}
