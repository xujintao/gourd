package service

import (
	"fmt"
	"log"

	// "github.com/gorilla/securecookie"
	"github.com/xujintao/gourd/apps/tpl/dao"
	"github.com/xujintao/gourd/apps/tpl/model"
)

// Users user list service
type users struct {
	db dao.DB
}

// GetUserList get user list
func (u *users) GetUserList() ([]*model.User, error) {
	users, err := u.db.GetUserList()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

// GetUser get the user by user name
func (u *users) GetUser(userName string) (*model.User, error) {
	user, err := u.db.GetUserByName(userName)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

// UpdateUser update user
func (u *users) UpdateUser(userName string, in *model.User) (*model.User, error) {
	user, err := u.db.GetUserByName(userName)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	user.Active = in.Active
	if err := u.db.UpdateUser(user); err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

// CreateUser create a user
func (u *users) CreateUser(in *model.User) (*model.User, error) {
	user := &model.User{
		Active: true,
		Name:   in.Name,
		Email:  in.Email,
		Avatar: in.Avatar,
		// Hash: base32.StdEncoding.EncodeToString(
		// 	securecookie.GenerateRandomKey(32),
		// ),
	}
	if err := user.Validate(); err != nil {
		log.Println(err)
		return nil, fmt.Errorf("user validate failed")
	}

	if err := u.db.CreateUser(user); err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

// DeleteUser delete a user
func (u *users) DeleteUser(userName string) error {
	user, err := u.db.GetUserByName(userName)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := u.db.DeleteUser(user); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
