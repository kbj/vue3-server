package v1

import "github.com/gin-gonic/gin"

type User struct {
	App *gin.RouterGroup
}

// Init 注册本文件内的路由
func (c User) Init() {
	c.App.GET("/:id", c.getUser)
}

// getUser 查询某个用户信息
func (c *User) getUser(context *gin.Context) {
	id := context.Param("id")
	context.JSON(200, map[string]interface{}{
		"id": id,
	})
}
