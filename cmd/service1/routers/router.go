package routers

import (
	"gorge/cmd/service1/controllers"
	"gorge/common/auth"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	//创建gin引擎
	router := gin.Default()

	router.GET("/ping", auth.JwtAuthMiddleware("hello") ,controllers.Pong)

	return router
}
