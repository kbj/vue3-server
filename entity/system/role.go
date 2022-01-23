package system

import (
	"vue3-server/entity/base"
)

type Role struct {
	base.Model
	Name  string `gorm:"comment:角色名" json:"name"`
	Intro string `gorm:"comment:角色介绍" json:"intro"`
}
