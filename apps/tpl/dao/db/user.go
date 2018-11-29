package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/xujintao/gorge/apps/tpl/dao/db/sqls"
	"github.com/xujintao/gorge/apps/tpl/model"
)

func (db *db) GetUserByID(id int64) (*model.User, error) {
	stmt := sqls.Lookup("user-find-id")
	var usr = new(model.User)

	if err := db.Get(usr, stmt, id); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
			return nil, fmt.Errorf("get user failed")
		}
	}
	return usr, nil
}

func (db *db) GetUserByName(login string) (*model.User, error) {
	stmt := sqls.Lookup("user-find-login")
	data := new(model.User)

	if err := db.Get(data, stmt, login); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
			return nil, fmt.Errorf("get user by login failed")
		}
	}
	return data, nil
}

func (db *db) GetUserList() ([]*model.User, error) {
	stmt := sqls.Lookup("user-find")
	data := []*model.User{}

	if err := db.Select(data, stmt); err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
			return nil, fmt.Errorf("get user list failed")
		}
	}
	return data, nil
}

func (db *db) GetUserCount() (count int, err error) {
	stmt := sqls.Lookup("count-users")

	err = db.Get(&count, stmt)
	return
}

func (db *db) CreateUser(user *model.User) error {
	stmt := sqls.Lookup("user-insert")

	if _, err := db.NamedExec(stmt, user); err != nil {
		log.Println(err)
		return fmt.Errorf("insert user failed")
	}
	return nil
}

func (db *db) UpdateUser(user *model.User) error {
	stmt := sqls.Lookup("user-update")

	if _, err := db.NamedExec(stmt, user); err != nil {
		log.Println(err)
		return fmt.Errorf("update user failed")
	}
	return nil
}

func (db *db) DeleteUser(user *model.User) error {
	stmt := sqls.Lookup("user-delete")

	_, err := db.Exec(stmt, user.ID)
	return err
}
