package response

import "vue3-server/entity/base"

type SysRoleListModel struct {
	base.Model
	Name     string              `json:"name"`
	Intro    string              `json:"intro"`
	MenuList *[]SysMenuRecursive `gorm:"-" json:"menuList"`
}
