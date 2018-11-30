package http

import (
	"github.com/gin-gonic/gin"
	"github.com/xujintao/gourd/apps/tpl/http/middleware/session"
)

// GetUser get user info
func GetUser(c *gin.Context) {
	c.JSON(200, session.GetUser(c))
}

func GetUserToken(c *gin.Context) {

}

func SyncRepoList(c *gin.Context) {

}

func GetRepoList(c *gin.Context) {

}
