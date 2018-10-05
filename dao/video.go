package dao

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //
	"github.com/jmoiron/sqlx"
	"github.com/xujintao/gorge/config"
	"github.com/xujintao/gorge/model"
)

// DB 抽象sqlx
var DB *sqlx.DB

func init() {
	var err error
	DB, err = sqlx.Connect("mysql", config.Config.DB.DSN)
	if err != nil {
		log.Fatalln(err)
	}

	DB.SetMaxOpenConns(config.Config.DB.MaxConn)

	go func() {
		if err := DB.Ping(); err != nil {
			log.Fatalln(err)
		}
		log.Println("db connected.")
		// 可以在这里建表
	}()
}

// NewVideo 上传视频
func NewVideo(project *model.Video) (int, error) {
	// 写库
	var SQLInsertVideo = `
INSERT INTO video (videoName, linkURL, des, ownerID, ownerName) 
VALUES (:videoName, :linkURL, :des, :ownerID, :ownerName)`

	// 开始事务
	tx, err := DB.Beginx()
	if err != nil {
		tx.Rollback()
		log.Panic(err)
	}

	res, err := tx.NamedExec(SQLInsertVideo, project)
	if err != nil {
		tx.Rollback()
		log.Panic(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Panic(err)
	}
	project.ID = int(lastID)

	// 事务提交
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Panic(err)
	}
	return project.ID, nil
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
