package service

import (
	"github.com/xujintao/gourd/apps/tpl/dao/db"
	"github.com/xujintao/gourd/apps/tpl/conf"
)

var (
	User  *user
	Users *users
)

func init() {
	db := db.New(conf.Config.DB.DSN, 100)
	User = &user{db}
	Users = &users{db}
}
