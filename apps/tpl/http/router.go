package http

import (
	"github.com/gin-gonic/gin"
	"github.com/xujintao/gourd/apps/tpl/http/middleware/session"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {

	e := gin.Default()
	e.Use(session.SetUser())

	// user
	user := e.Group("/api/user", session.MustUser())
	{
		user.GET("", GetSelf)
		user.GET("/token", GetUserToken)

		user.POST("/repos", SyncRepoList)
		user.GET("/repos", GetRepoList)
	}

	// users
	users := e.Group("/api/users", session.MustAdmin())
	{
		users.GET("", GetUserList)
		users.GET("/:name", GetUser)
		users.POST("", CreateUser)
		users.PUT("/:name", UpdateUser)
		users.DELETE("/:name", DeleteUser)
	}

	// repo
	repo := e.Group("/api/repos/:group/:project")
	{
		repo.GET("", GetRepo)
		repo.GET("/builds", GetRepoBuildList)
	}

	// stream
	{
		// api.GET("/stream", GetStream)
	}

	return e
}
