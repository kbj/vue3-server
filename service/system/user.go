package system

import (
	"github.com/gofiber/fiber/v2"
	"vue3-server/common/global"
	"vue3-server/ent"
	"vue3-server/ent/user"
	"vue3-server/model/base"
	"vue3-server/model/system/request"
	"vue3-server/utils"
)

type UserService struct{}

// UserLogin 用户登录
func (userService *UserService) UserLogin(ctx *fiber.Ctx, userInfo *request.SysUserModel) *base.ResponseEntity {
	u, err := global.Db.User.Query().
		Where(user.Username(userInfo.Name), user.Password(userInfo.Password)).Only(ctx.UserContext())
	if err != nil && ent.IsNotFound(err) {
		return utils.ResponseFail("请检查用户名密码是否正确！")
	}

	// 生成token
	claims := base.BaseClaims{
		ID:       u.ID,
		Username: u.Username,
	}
	token, err := utils.CreateJwtToken(utils.CreateClaims(claims))
	if err != nil {
		global.Logger.Error("创建token失败，失败原因：" + err.Error())
		return utils.ResponseFail("系统错误！")
	}

	result := map[string]interface{}{
		"id":    u.ID,
		"name":  u.Username,
		"token": token,
	}

	return utils.ResponseSuccess(result)
}
