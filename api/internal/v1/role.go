package v1

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"vue3-server/model/system/request"
	"vue3-server/utils"
)

func InitRoleRoute(route *fiber.Router) {
	router := *route
	router.Get("/:roleId/menu", getRoleMenus)
	router.Post("/list", getRoleList)
	router.Post("", createRole)
	router.Patch("/:roleId", updateRole)
	router.Delete("/:roleId", deleteRole)
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

	return context.JSON(menuService.RoleMenus(uint(atoi)))
}

// 查询角色列表
func getRoleList(context *fiber.Ctx) error {
	var param request.SysRoleListModel
	err := context.BodyParser(&param)
	if err != nil {
		return context.JSON(utils.ResponseFail("参数解析失败！" + err.Error()))
	}
	return context.JSON(roleService.GetRoleList(&param))
}

// 创建角色
func createRole(context *fiber.Ctx) error {
	var param request.SysRoleMenuModel
	if err := context.BodyParser(&param); err != nil {
		return context.JSON(utils.ResponseFail("参数解析失败！"))
	}
	return context.JSON(roleService.CreateRole(&param))
}

// 更新角色
func updateRole(context *fiber.Ctx) error {
	roleId, err := strconv.Atoi(context.Params("roleId"))
	if err != nil {
		return context.JSON(utils.ResponseFail("获取角色ID失败！"))
	}
	var param request.SysRoleMenuModel
	if err := context.BodyParser(&param); err != nil {
		return context.JSON(utils.ResponseFail("参数解析失败！"))
	}
	param.Id = uint(roleId)

	return context.JSON(roleService.UpdateRole(&param))
}

// 删除角色
func deleteRole(context *fiber.Ctx) error {
	roleId, err := strconv.Atoi(context.Params("roleId"))
	if err != nil {
		return context.JSON(utils.ResponseFail("获取角色ID失败！"))
	}
	return context.JSON(roleService.DeleteRole(uint(roleId)))
}
