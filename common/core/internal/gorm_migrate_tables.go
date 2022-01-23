package internal

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"vue3-server/common/global"
	"vue3-server/entity/system"
)

// InitializeTables 自动初始化好数据库表
func InitializeTables(db *gorm.DB) {
	err := db.AutoMigrate(
		system.User{},
		system.Role{},
		system.Department{},
		system.Menu{},
		system.RoleMenu{},
	)

	if err != nil {
		global.Logger.Error("初始化表结构失败！", zap.Error(err))
		os.Exit(0)
	}
	global.Logger.Info("初始化表结构成功！")
}
