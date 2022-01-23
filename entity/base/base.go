package base

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID       uint           `gorm:"comment:主键;primarykey" json:"id"`
	CreateAt time.Time      `gorm:"comment:创建时间" json:"createAt"`
	UpdateAt time.Time      `gorm:"comment:修改时间" json:"updateAt"`
	DeleteAt gorm.DeletedAt `gorm:"comment:删除时间;index" json:"deleteAt"`
}
