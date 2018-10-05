package model

import "time"

// Video 项目
type Video struct {
	ID         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"videoName"`
	URL        string    `json:"linkURL" db:"linkURL"`
	Des        string    `json:"des" db:"des"`
	OwnerID    int       `json:"ownerID" db:"ownerID"`
	OwnerName  string    `json:"ownerName" db:"ownerName"`
	Delete     int       `db:"deleteStatus"`
	CreateTime time.Time `db:"createTime"`
	UpdateTime time.Time `db:"updateTime"`
}
