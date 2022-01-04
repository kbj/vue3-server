package v1

import (
	"github.com/gofiber/fiber/v2"
	"vue3-server/utils"
)

func InitUserRoute(route *fiber.Router) {
	router := *route
	router.Get("/:id", getUser)
}

// getUser 查询某个用户信息
func getUser(context *fiber.Ctx) error {
	id := context.Params("id")
	return context.JSON(utils.ResponseSuccess(id))
}
