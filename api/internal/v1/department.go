package v1

import (
	"github.com/gofiber/fiber/v2"
	"vue3-server/model/system/request"
	"vue3-server/utils"
)

func InitDepartmentRoute(route *fiber.Router) {
	router := *route
	router.Post("/list", getDepartmentList)
}

// 部门列表
func getDepartmentList(context *fiber.Ctx) error {
	var param request.SysUserListModel
	err := context.BodyParser(&param)
	if err != nil {
		return context.JSON(utils.ResponseFail("解析参数失败！" + err.Error()))
	}
	return context.JSON(departmentService.GetDepartmentList(&param))
}
