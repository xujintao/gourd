package service

import (
	"log"

	"github.com/xujintao/gourd/apps/tpl/dao"
	"github.com/xujintao/gourd/apps/tpl/lib/token"
	"github.com/xujintao/gourd/apps/tpl/model"
)

// User current user
type user struct {
	dao dao.Dao
}

func (u *user) GetUserToken(user *model.User) (string, error) {
	claims := token.New(token.UserToken, user.Name)
	raw, err := claims.Sign(user.Hash)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return raw, nil
}

// GetFeedList get feed list
func (u *user) GetFeedList(userName string, latest bool) ([]*model.Feed, error) {
	user, err := u.dao.GetUserByName(userName)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// sync here

	if latest {
		feeds, err := u.dao.GetFeedListLatest(user)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return feeds, nil
	}

	feeds, err := u.dao.GetFeedList(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return feeds, nil
}

// GetRepoList get repo list
func (u *user) GetRepoList(user *model.User, all, flush bool) ([]*model.Repo, error) {

	// sync here

	repos, err := u.dao.GetRepoList(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if all {
		return repos, nil
	}

	active := []*model.Repo{}
	for _, repo := range repos {
		if repo.IsActive {
			active = append(active, repo)
		}
	}

	return active, nil
}
