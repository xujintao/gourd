package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xujintao/gourd/apps/tpl/http/middleware/session"
	"github.com/xujintao/gourd/apps/tpl/service"
)

// GetUser get user info
func GetUser(c *gin.Context) {
	c.JSON(200, session.GetUser(c))
}

// GetUserToken get user info with personal token
func GetUserToken(c *gin.Context) {
	user := session.GetUser(c)
	raw, err := service.User.GetUserToken(user)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.String(200, raw)
}

func SyncRepoList(c *gin.Context) {

}

func GetRepoList(c *gin.Context) {

}
