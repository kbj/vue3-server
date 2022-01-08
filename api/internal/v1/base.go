package v1

import (
	"github.com/gofiber/fiber/v2"
	"vue3-server/utils"
)

// InitBaseRoute 无需验证登录权限的路由
func InitBaseRoute(route *fiber.Router) {
	router := *route
	router.Get("/health", checkHealth)
	router.Post("/login", userLogin)
}

// 健康检查
func checkHealth(context *fiber.Ctx) error {
	return context.JSON(utils.ResponseSuccess("ok"))
}

// 用户登录
func userLogin(context *fiber.Ctx) error {
	// 获取登录信息
	username := context.FormValue("username")
	password := context.FormValue("password")

	var param = make(map[string]string)
	param["username"] = username
	param["password"] = password

	return context.JSON(param)
}
