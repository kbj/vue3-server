package system

import "vue3-server/entity/base"

type Menu struct {
	base.Model
	Name       string `gorm:"comment:名称" json:"name"`
	Url        string `gorm:"comment:地址" json:"url"`
	Icon       string `gorm:"comment:图标" json:"icon"`
	Sort       uint   `gorm:"comment:排序" json:"sort"`
	ParentId   uint   `gorm:"comment:父级ID" json:"parentId"`
	Type       uint8  `gorm:"comment:类型" json:"type"`
	Permission string `json:"permission"`
}
