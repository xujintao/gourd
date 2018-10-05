package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xujintao/gorge/model"
	"github.com/xujintao/gorge/service"
)

// NewVideo 上传视频
func NewVideo(c *gin.Context) {
	video := &model.Video{}
	if err := c.ShouldBind(video); err != nil {
		c.JSON(200, gin.H{
			"code":    10101,
			"message": fmt.Errorf("参数解析失败%s", err.Error()),
		})
		return
	}

	vid, err := service.NewVideo(video)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    10102,
			"message": fmt.Errorf("上传视频失败%s", err.Error()),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": vid,
	})
	return
}

// GetVideos 获取项目列表
func GetVideos(c *gin.Context) {

	// 获取uid
	v, _ := c.Get("uid")
	uid, _ := v.(string)

	// 调用
	videos, err := service.GetVideos(uid)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    10201,
			"message": fmt.Errorf("获取项目失败%s", err.Error()),
		})
		return
	}

	// 成功返回
	c.JSON(200, gin.H{
		"code": 200,
		"data": videos,
	})
}

// GetVideo 获取视频详情
func GetVideo(c *gin.Context) {

	// 获取vid
	vid := c.Param("vid")

	// 调用
	video, err := service.GetVideo(vid)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    10201,
			"message": fmt.Errorf("获取项目失败%s", err.Error()),
		})
		return
	}

	// 成功返回
	c.JSON(200, gin.H{
		"code": 200,
		"data": video,
	})
}
