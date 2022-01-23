package v1

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"vue3-server/utils"
)

func InitRoleRoute(route *fiber.Router) {
	router := *route
	router.Get("/:roleId/menu", getRoleMenus)
}

// 查询对应角色所拥有的菜单
func getRoleMenus(context *fiber.Ctx) error {
	roleId := context.Params("roleId")
	if roleId == "" {
		return context.JSON(utils.ResponseFail("请传入相关的角色ID！"))
	}
	atoi, err := strconv.Atoi(roleId)
	if err != nil {
		return context.JSON(utils.ResponseFail("请传入正确的角色ID！"))
	}

	return context.JSON(menuService.RoleMenus(context, uint(atoi)))
}
