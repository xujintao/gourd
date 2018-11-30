package session

import (
	"github.com/gin-gonic/gin"
	"github.com/xujintao/gourd/apps/tpl/lib/token"
	"github.com/xujintao/gourd/apps/tpl/model"
	"github.com/xujintao/gourd/apps/tpl/service"
)

// GetUser get user from gin context
func GetUser(c *gin.Context) *model.User {
	v, ok := c.Get("user")
	if !ok {
		return nil
	}

	u, ok := v.(*model.User)
	if !ok {
		return nil
	}
	return u
}

// SetUser set user to gin context
// 1, extract token from session
// 2, parse token
// 3, get user from service
// 4, set user to gin context
func SetUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		var user *model.User

		claims, err := token.ParseRequest(c.Request, func(userName string) (string, error) {
			var err error
			user, err = service.Users.GetUser(userName)
			return user.Hash, err
		})
		if err == nil {
			confv := c.MustGet("config")
			if conf, ok := confv.(*model.Settings); ok {
				user.Admin = conf.IsAdmin(user)
			}
			c.Set("user", user)

			// if this is a session token (ie not the API token)
			// this means the user is accessing with a web browser,
			// so we should implement CSRF protection measures.
			if claims.Kind == token.SessionToken {
				err := token.CheckCSRF(c.Request, func(userName string) (string, error) {
					return user.Hash, nil
				})
				// if csrf token validation fails, exit immediately
				// with a not authorized error.
				if err != nil {
					c.AbortWithStatus(401)
				}
			}
		}

		c.Next()
	}
}

// MustUser must be user
func MustUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		user := GetUser(c)
		if user == nil {
			c.AbortWithStatus(401)
		}

		c.Next()
	}
}
