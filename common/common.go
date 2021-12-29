package common

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"vue3-server/api"
)

// Init 初始化配置
func Init(app *fiber.App) {
	// 初始化日志
	initLog(app)
	// 注册路由
	api.RegisterRoute(app)
}

// Start 启动服务
func Start(app *fiber.App, listen string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(listen); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")
}

// 初始化日志
func initLog(app *fiber.App) {

}
