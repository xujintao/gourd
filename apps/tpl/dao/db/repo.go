package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/xujintao/gorge/apps/tpl/dao/db/sqls"
	"github.com/xujintao/gorge/apps/tpl/model"
)

func (db *db) GetFeedList(user *model.User) ([]*model.Feed, error) {
	stmt := sqls.Lookup("feed")
	data := []*model.Feed{}

	if err := db.Select(&data, stmt, user.ID); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
			return nil, fmt.Errorf("get feed list failed")
		}
	}
	return data, nil
}

func (db *db) GetFeedListLatest(user *model.User) ([]*model.Feed, error) {
	stmt := sqls.Lookup("feed-latest-build")
	data := []*model.Feed{}

	if err := db.Select(&data, stmt, user.ID); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
			return nil, fmt.Errorf("get feed list latest failed")
		}
	}
	return data, nil
}

func (db *db)GetRepoList(user *model.User) ([]*model.Repo, error){
	return nil,nil
}
