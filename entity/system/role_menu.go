package system

import "vue3-server/entity/base"

type RoleMenu struct {
	base.Model
	RoleId uint `gorm:"comment:角色ID" json:"roleId"`
	Role   Role `gorm:"foreignKey:RoleId" json:"-"`
	MenuId uint `gorm:"comment:菜单ID" json:"menuId"`
	Menu   Menu `gorm:"foreignKey:MenuId" json:"-"`
}
