package v1

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"vue3-server/utils"
)

func InitUserRoute(route *fiber.Router) {
	router := *route
	router.Get("/:id", getUser)
}

// getUser 查询某个用户信息
func getUser(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		return context.JSON(utils.ResponseFail("请选择要查询的用户信息！"))
	}
	newId, err := strconv.Atoi(id)
	if err != nil {
		return context.JSON(utils.ResponseFail("ID错误！"))
	}

	return context.JSON(userService.GetUserInfo(context, uint(newId)))
}
