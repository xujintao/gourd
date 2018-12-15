package http

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xujintao/gourd/apps/tpl/http/middleware/session"
	"github.com/xujintao/gourd/apps/tpl/service"
)

// GetSelf get user info
func GetSelf(c *gin.Context) {
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

// SyncRepoList synchronize lated repository from gitlab
func SyncRepoList(c *gin.Context) {

}

// GetRepoList get user related repo list
func GetRepoList(c *gin.Context) {
	var (
		user     = session.GetUser(c)
		all, _   = strconv.ParseBool(c.Query("all"))
		flush, _ = strconv.ParseBool(c.Query("flush"))
	)

	repos, err := service.User.GetRepoList(user, all, flush)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, repos)
}
