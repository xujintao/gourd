package service

import (
	"log"

	gitlab "github.com/xanzy/go-gitlab"
	"github.com/xujintao/gourd/apps/tpl/dao"
	"github.com/xujintao/gourd/apps/tpl/model"
)

type repo struct {
	db           dao.DB
	gitlibClient *gitlab.Client
}

func (repo *repo) GetRepo(group, project string) (*model.Repo, error) {
	name := group + "/" + project
	r, err := repo.db.GetRepoByName(name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return r, nil
}
