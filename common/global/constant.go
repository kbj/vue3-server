package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"vue3-server/common/config"
)

var (
	Logger *zap.Logger   // 日志组件
	VP     *viper.Viper  // 配置对象
	Config config.Server // 配置文件
)

var (
	// ConfigEnv 配置文件的环境变量名
	ConfigEnv = "GVA_CONFIG"

	// ConfigFile 配置文件的路径
	ConfigFile = "config.yaml"
)
