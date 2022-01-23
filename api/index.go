package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	v1 "vue3-server/api/internal/v1"
	"vue3-server/common/global"
	"vue3-server/common/middleware"
	"vue3-server/utils"
)

// RegisterRoute 注册系统所有的路由
func RegisterRoute(app *fiber.App) {
	registerAnonymousRoutes(app)
	registerAuthRoutes(app)
	register404Routes(app)
}

// 注册匿名无需登录校验的路由
func registerAnonymousRoutes(app *fiber.App) {
	// 允许跨域注册路由
	annoyGroup := app.Group("").Use(cors.New())
	v1.InitBaseRoute(&annoyGroup)
	global.Logger.Info("初始化匿名路由！")
}

// 注册需要登录验证的路由
func registerAuthRoutes(app *fiber.App) {
	v1Group := app.Group("")
	v1Group.Use(middleware.AuthLogin()).Use(cors.New())
	global.Logger.Info("添加中间件：AuthLogin")

	// 用户相关
	userGroup := v1Group.Group("users")
	v1.InitUserRoute(&userGroup)

	// 角色相关
	roleGroup := v1Group.Group("role")
	v1.InitRoleRoute(&roleGroup)
	global.Logger.Info("初始化路由完成！")
}

// 注册404路由，兜底用
func register404Routes(app *fiber.App) {
	app.Use("/**", func(c *fiber.Ctx) error {
		c.Status(fiber.StatusNotFound)
		return c.JSON(utils.ResponseNotFound())
	})
}
