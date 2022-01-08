package core

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"vue3-server/common/global"
)

// InitializeSession 初始化好一个session池，默认先存内存，后续可以改放进redis
func InitializeSession() *session.Store {
	global.Logger.Info("初始化session存储")
	return session.New()
}
