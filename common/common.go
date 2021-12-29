package common

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"vue3-server/api"
	"vue3-server/entity"
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

// ErrorHandler 通用的错误处理逻辑
func ErrorHandler() func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		errorCode := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			errorCode = e.Code
		}

		// 全局使用JSON方式返回错误
		return c.Status(errorCode).JSON(entity.ResponseEntity{
			Code: 500,
			Data: err.Error(),
		})
	}
}

// 初始化日志
func initLog(app *fiber.App) {

}
