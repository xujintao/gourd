package service

import (
	"github.com/xujintao/gorge/apps/video/dao"
	"github.com/xujintao/gorge/apps/video/model"
)

// NewVideo 上传视频
func NewVideo(video *model.Video) (int, error) {
	vid, err := dao.NewVideo(video)
	if err != nil {
		return 0, err
	}
	return vid, nil
}

// GetVideos 获取所有视频
func GetVideos(uid string) ([]model.Video, error) {
	videos, err := dao.GetVideos(uid)
	if err != nil {
		return nil, err
	}

	return videos, nil
}

// GetVideo 获取视频详情
func GetVideo(vid string) (*model.Video, error) {
	video, err := dao.GetVideo(vid)
	if err != nil {
		return nil, err
	}

	return video, nil
}
