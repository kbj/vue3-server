package v1

import (
	"github.com/gofiber/fiber/v2"
	"vue3-server/utils"
)

// InitBaseRoute 无需验证登录权限的路由
func InitBaseRoute(route *fiber.Router) {
	router := *route
	router.Get("/health", checkHealth)
}

// 健康检查
func checkHealth(context *fiber.Ctx) error {
	return context.JSON(utils.ResponseSuccess("ok"))
}
