package dao

import (
	"database/sql"
	"log"

	"github.com/xujintao/gorge/apps/video/model"
)

// NewVideo 上传视频
func NewVideo(video *model.Video) (int, error) {
	// 开始事务
	tx, err := DB.Beginx()
	if err != nil {
		tx.Rollback()
		log.Panic(err)
	}

	// 写库
	var SQLInsertVideo = `
		INSERT INTO video (videoName, linkURL, des, ownerID, ownerName) 
		VALUES (:videoName, :linkURL, :des, :ownerID, :ownerName)`
	res, err := tx.NamedExec(SQLInsertVideo, video)
	if err != nil {
		tx.Rollback()
		log.Panic(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Panic(err)
	}
	video.ID = int(lastID)

	// 事务提交
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Panic(err)
	}
	return video.ID, nil
}

// GetVideos 获取所有视频
func GetVideos(uid string) ([]model.Video, error) {

	// 按照道理来说是获取uid关联出来的所有视频，这里只是演示一下

	var SQLGetVideos = `SELECT * FROM video WHERE deleteStatus = 0`

	// 读库
	videos := make([]model.Video, 0)
	if err := DB.Select(&videos, SQLGetVideos); err != nil {
		if err != sql.ErrNoRows {
			log.Panic(err)
		}
	}
	return videos, nil
}

// GetVideo 获取视频详情
func GetVideo(vid string) (*model.Video, error) {

	var SQLGetVideo = `SELECT * FROM video WHERE id = ? AND deleteStatus = 0`

	// 读库
	video := &model.Video{}
	if err := DB.Get(video, SQLGetVideo, vid); err != nil {
		if err != sql.ErrNoRows {
			log.Panic(err)
		}
	}

	return video, nil
}
