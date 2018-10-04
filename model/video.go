package model

import "time"

// Video 项目
type Video struct {
	ID         int       `db:"id" json:"id"`
	Name       string    `db:"videoName" json:"name"`
	URL        string    `db:"linkURL" json:"linkURL"`
	Des        string    `db:"des" json:"des"`
	OwnerID    int       `db:"ownerID" json:"ownerID"`
	OwnerName  string    `db:"ownerName" json:"ownerName"`
	Delete     int       `db:"deleteStatus" json:"-"`
	CreateTime time.Time `db:"createTime" json:"createTime"`
	UpdateTime time.Time `db:"updateTime" json:"updateTime"`
}
