package system

import "vue3-server/entity/base"

type Department struct {
	base.Model
	Name     string `gorm:"comment:部门名称" json:"name"`
	Leader   string `gorm:"comment:领导名称" json:"leader"`
	ParentId uint   `gorm:"comment:父ID" json:"parentId"`
}
