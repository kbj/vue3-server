package system

import (
	"vue3-server/entity/base"
)

type User struct {
	base.Model
	Name         string     `gorm:"index;comment:姓名;<-:create" json:"name"`
	Realname     string     `gorm:"comment:真实姓名" json:"realname"`
	Password     string     `gorm:"comment:密码" json:"-"`
	Cellphone    string     `gorm:"comment:手机号;size:11" json:"cellphone"`
	Enable       uint       `gorm:"comment:状态 1是启用0是禁用;size:1" json:"enable"`
	Role         Role       `gorm:"foreignKey:RoleId" json:"role"`
	RoleId       uint       `gorm:"comment:角色ID" json:"-"`
	Department   Department `gorm:"foreignKey:DepartmentId" json:"department"`
	DepartmentId uint       `gorm:"comment:部门ID" json:"-"`
}
