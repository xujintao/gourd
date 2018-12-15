package http

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xujintao/gourd/apps/tpl/model"
	"github.com/xujintao/gourd/apps/tpl/service"
)

// GetUserList 获取用户列表
func GetUserList(c *gin.Context) {
	users, err := service.Users.GetUserList()
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code":    20101,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": users,
	})
}

// GetUser 获取用户
func GetUser(c *gin.Context) {
	user, err := service.Users.GetUser(c.Param("name"))
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code":    20201,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": user,
	})
}

// CreateUser 创建用户
func CreateUser(c *gin.Context) {
	// 获取表单
	in := model.User{}
	if err := c.ShouldBind(&in); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code":    20301,
			"message": fmt.Sprintf("bind user failed, %#v", in),
		})
		return
	}

	// 验证表单
	if err := in.Validate(); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code":    20302,
			"message": err,
		})
		return
	}

	// 创建用户
	u, err := service.Users.CreateUser(&in)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code":    20303,
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": u,
	})
}

// UpdateUser 更新用户
func UpdateUser(c *gin.Context) {
	// 获取表单
	in := model.User{}
	if err := c.ShouldBind(&in); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code":    20401,
			"message": fmt.Sprintf("bind user failed, %#v", in),
		})
		return
	}

	// 表单验证
	if err := in.Validate(); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code":    20402,
			"message": err,
		})
		return
	}

	// 修改用户profile
	u, err := service.Users.UpdateUser(in.Name, &in)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code":    20403,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": u,
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	name := c.Param("name")

	// 这里为了验证name强行构造了一个User结构体
	// 是否有其他好方法
	user := model.User{
		Name: name,
	}
	if err := user.Validate(); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code":    20501,
			"message": err.Error(),
		})
		return
	}

	// 删除用户
	err := service.Users.DeleteUser(name)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code":    20502,
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
	})
}
