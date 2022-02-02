package system

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"time"
	"vue3-server/common/global"
	"vue3-server/entity/system"
	"vue3-server/model/base"
	"vue3-server/model/system/request"
	"vue3-server/model/system/response"
	"vue3-server/utils"
)

type UserService struct{}

// UserLogin 用户登录
func (userService *UserService) UserLogin(ctx *fiber.Ctx, userInfo *request.SysUserModel) *base.ResponseEntity {
	// 将密码加密
	password := utils.Md5Encode(userInfo.Password)

	var user system.User
	if err := global.Db.Where("name = ? and password = ?", userInfo.Name, password).First(&user).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.ResponseFail("您的用户名或密码错误！")
	}

	// 生成token
	claims := base.BaseClaims{
		ID:       user.ID,
		Username: user.Name,
	}
	token, err := utils.CreateJwtToken(utils.CreateClaims(claims))
	if err != nil {
		global.Logger.Error("创建token失败，失败原因：", zap.Error(err))
		return utils.ResponseFail("系统错误！")
	}

	result := map[string]interface{}{
		"id":    user.ID,
		"name":  user.Name,
		"token": token,
	}

	return utils.ResponseSuccess(result)
}

// GetUserInfo 查询用户信息
func (userService *UserService) GetUserInfo(ctx *fiber.Ctx, id uint) *base.ResponseEntity {
	var user system.User
	user.ID = id

	// 用户基本信息
	err := global.Db.Model(&user).Preload("Role").Preload("Department").First(&user).Error
	if err != nil {
		return utils.ResponseFail("用户不存在!")
	}

	// 封装
	resp := response.SysUserModel{
		ID:         user.ID,
		Name:       user.Name,
		Realname:   user.Realname,
		Cellphone:  user.Cellphone,
		Enable:     user.Enable,
		Role:       user.Role,
		Department: user.Department,
	}
	return utils.ResponseSuccess(&resp)
}

// GetUserList 用户列表
func (userService UserService) GetUserList(ctx *fiber.Ctx, param *request.SysUserListModel) *base.ResponseEntity {
	db := global.Db.Model(&system.User{}).Offset(param.Offset).Limit(param.Size)
	if param.Name != nil && *param.Name != "" {
		db = db.Where("name like ?", fmt.Sprintf("%%%s%%", *param.Name))
	}
	if param.Realname != nil && *param.Realname != "" {
		db = db.Where("realname like ?", fmt.Sprintf("%%%s%%", *param.Realname))
	}
	if param.Id != "" {
		db = db.Where("id = ?", param.Id)
	}
	if param.Cellphone != nil && *param.Cellphone != "" {
		db = db.Where("cellphone = ?", *param.Cellphone)
	}
	if param.Enable != nil {
		switch param.Enable.(type) {
		case float64:
			db = db.Where("enable = ?", param.Enable)
		}
	}
	if param.CreateAt != nil {
		switch reflect.TypeOf(param.CreateAt).Kind() {
		case reflect.Array, reflect.Slice:
			// 时间数组
			values := reflect.ValueOf(param.CreateAt)
			startString := values.Index(0).Interface().(string)
			endString := values.Index(1).Interface().(string)
			start, _ := time.Parse("2006-01-02T15:04:05.000Z", startString)
			end, _ := time.Parse("2006-01-02T15:04:05.000Z", endString)
			db.Where("create_at between ? and ?", start, end)
		}
	}

	var lists []response.SysUserListModel
	var total int64
	db.Count(&total)
	db.Find(&lists)

	var result = make(map[string]interface{})
	result["list"] = lists
	result["totalCount"] = total
	return utils.ResponseSuccess(&result)
}

// UpdateUserInfo 更新用户信息
func (*UserService) UpdateUserInfo(ctx *fiber.Ctx, param *request.SysUser) *base.ResponseEntity {
	var user system.User
	global.Db.First(&user, param.Id)
	global.Db.Model(&user).Updates(system.User{
		Cellphone:    param.Cellphone,
		DepartmentId: param.DepartmentId,
		Name:         param.Name,
		Realname:     param.Realname,
		RoleId:       param.RoleId,
	})
	return utils.ResponseSuccess("更新成功！")
}
