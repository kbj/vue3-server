package boot

import (
	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"vue3-server/api"
	"vue3-server/common/core"
	"vue3-server/common/global"
	"vue3-server/model/system"
)

// Init 初始化配置
func Init(app *fiber.App) {
	// 初始化配置文件
	global.VP = core.Viper()

	// 初始化Zap日志框架
	global.Logger = core.InitializeZap()

	// 初始化数据库
	if global.Db = core.InitializeEntInstance(); global.Db != nil {
		defer global.Db.Close()
	}

	// 初始化session池
	global.Session = core.InitializeSession()

	// fiber框架的日志改为zap
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: global.Logger,
	}))

	// 注册路由
	api.RegisterRoute(app)
}

// Start 启动服务
func Start(app *fiber.App, listen string) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		global.Logger.Info("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(listen); err != nil {
		global.Logger.Error(err.Error())
	}

	global.Logger.Info("Running cleanup tasks...")
}

// ErrorHandler 通用的错误处理逻辑
func ErrorHandler() func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		errorCode := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			errorCode = e.Code
		}

		// 全局使用JSON方式返回错误
		return c.Status(errorCode).JSON(system.ResponseEntity{
			Code: 500,
			Data: err.Error(),
		})
	}
}
