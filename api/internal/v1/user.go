package v1

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
	"vue3-server/common/global"
	"vue3-server/model/system/request"
	"vue3-server/utils"
)

func InitUserRoute(route *fiber.Router) {
	router := *route
	router.Get("/:id", getUser)
	router.Post("/list", userList)
	router.Patch("/:id", updateUserInfo)
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

// 用户列表
func userList(context *fiber.Ctx) error {
	var param request.SysUserListModel
	err := context.BodyParser(&param)
	if err != nil {
		global.Logger.Error("解析用户列表解析参数失败！", zap.Error(err))
		return context.JSON(utils.ResponseFail(fmt.Sprintf("解析用户列表解析参数失败！%v", err.Error())))
	}

	return context.JSON(userService.GetUserList(context, &param))
}

// 更新用户信息
func updateUserInfo(context *fiber.Ctx) error {
	id, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		return context.JSON(utils.ResponseFail("请输入正常的ID！"))
	}
	var param request.SysUser
	err = context.BodyParser(&param)
	if err != nil {
		return context.JSON(utils.ResponseFail("解析参数失败！"))
	}
	param.Id = uint(id)

	return context.JSON(userService.UpdateUserInfo(context, &param))
}
