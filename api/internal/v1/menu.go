package v1

import (
	"github.com/gofiber/fiber/v2"
	"vue3-server/model/system/request"
	"vue3-server/utils"
)

func InitMenuRoute(route *fiber.Router) {
	router := *route
	router.Post("/list", getMenuList)
}

// 查询菜单列表
func getMenuList(context *fiber.Ctx) error {
	var param request.SysUserListModel
	if err := context.BodyParser(&param); err != nil {
		return context.JSON(utils.ResponseFail("解析参数失败！" + err.Error()))
	}
	return context.JSON(menuService.GetMenuList(param.Offset, param.Size))
}
