package global

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"vue3-server/common/config"
)

var (
	Logger             *zap.Logger             // 日志组件
	VP                 *viper.Viper            // 配置对象
	Config             config.Server           // 配置文件
	Session            *session.Store          // 全局session池
	ConcurrencyControl = &singleflight.Group{} // 并发控制
	Db                 *gorm.DB                // 数据库Orm对象
)

var (
	// ConfigEnvName 配置文件的环境变量名
	ConfigEnvName = "GVA_CONFIG"

	// ConfigFileName 配置文件的路径
	ConfigFileName = "config.yaml"
)
