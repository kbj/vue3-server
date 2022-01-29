package response

import (
	"vue3-server/entity/base"
	"vue3-server/entity/system"
)

type SysUserModel struct {
	ID         uint              `json:"id"`
	Name       string            `json:"name"`
	Realname   string            `json:"realname"`
	Cellphone  string            `json:"cellphone"`
	Enable     uint              `json:"enable"`
	Role       system.Role       `json:"role"`
	Department system.Department `json:"department"`
}

type SysUserListModel struct {
	base.Model
	Name         string `json:"name"`
	Realname     string `json:"realname"`
	Cellphone    string `json:"cellphone"`
	Enable       uint   `json:"enable"`
	DepartmentId uint   `json:"departmentId"`
	RoleId       uint   `json:"roleId"`
}
