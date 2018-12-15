package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xujintao/gourd/apps/tpl/service"
)

// GetRepo 获取仓库详情
func GetRepo(c *gin.Context) {
	group := c.Param("group")
	project := c.Param("project")

	// validate

	repo, err := service.Repo.GetRepo(group, project)
	if err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"code":    30101,
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": repo,
	})
}

// GetRepoBuildList 获取仓库构建列表
func GetRepoBuildList(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "deving",
	})
}
