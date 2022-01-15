package system

import (
	"github.com/gofiber/fiber/v2"
	"vue3-server/model/base"
	"vue3-server/model/system/request"
)

type UserService struct{}

// UserLogin 用户登录
func (userService *UserService) UserLogin(ctx *fiber.Ctx, userInfo *request.SysUserModel) *base.ResponseEntity {
	return nil
}
