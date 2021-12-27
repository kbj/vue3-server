package api

import (
	"github.com/gin-gonic/gin"
	"vue3-server/api/internal/v1"
)

// RegisterRoute 注册系统所有的路由
func RegisterRoute(app *gin.Engine) {
	// user部分的路由
	v1.User{App: app.Group("/v1/user")}.Init()
}
