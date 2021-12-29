package v1

import (
	"github.com/gofiber/fiber/v2"
	"vue3-server/utils"
)

type User struct {
	App fiber.Router
}

// Init 注册本文件内的路由
func (c User) Init() {
	c.App.Get("/:id", c.getUser)
}

// getUser 查询某个用户信息
func (c *User) getUser(context *fiber.Ctx) error {
	id := context.Params("id")
	return context.JSON(utils.ResponseSuccess(id))
}
