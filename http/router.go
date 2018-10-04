package http

import (
	"github.com/gin-gonic/gin"
	"github.com/xujintao/gorge/model"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	//创建gin引擎
	router := gin.Default()
	router.Use(JwtAuthMiddleware(model.Config.Auth.JWT.Secret))

	router.POST("/videos", NewVideo)
	router.GET("/videos", GetVideos)
	router.GET("/videos/:vid", GetVideo)

	return router
}
