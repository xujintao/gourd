package service

import (
	"github.com/xujintao/gorge/apps/tpl/dao/db"
)

var (
	User  *user
	Users *users
)

func init() {
	db := db.New("root:1234@tcp(127.0.0.1:3306)/gourd", 100)
	User = &user{db}
	Users = &users{db}
}
