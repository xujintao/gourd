package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/xujintao/gourd/apps/tpl/dao/db/sqls"
	"github.com/xujintao/gourd/apps/tpl/model"
)

func (db *db) GetRepoByID(id int) (*model.Repo, error) {
	stmt := sqls.Lookup("repo-find-id")
	var repo = new(model.Repo)

	if err := db.Get(repo, stmt, id); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
			return nil, fmt.Errorf("get repo by id failed")
		}
	}
	return repo, nil
}

func (db *db) GetRepoByName(name string) (*model.Repo, error) {
	stmt := sqls.Lookup("repo-find-name")
	var repo = new(model.Repo)

	if err := db.Get(repo, stmt, name); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
			return nil, fmt.Errorf("get repo by name failed")
		}
	}
	return repo, nil
}

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

func (db *db) GetRepoList(user *model.User) ([]model.Repo, error) {
	stmt := sqls.Lookup("repo-find-user")
	data := []model.Repo{}

	if err := db.Select(&data, stmt, user); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
			return nil, fmt.Errorf("get repo list failed")
		}
	}

	return data, nil
}
