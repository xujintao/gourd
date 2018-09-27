package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xujintao/gorge/cmd/service1/controllers"
	"github.com/xujintao/gorge/common/auth"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	//创建gin引擎
	router := gin.Default()

	router.GET("/ping", auth.JwtAuthMiddleware("hello"), controllers.Pong)

	return router
}
