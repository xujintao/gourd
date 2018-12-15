package service

import (
	"net/http"
	"time"

	gitlab "github.com/xanzy/go-gitlab"
	"github.com/xujintao/gourd/apps/tpl/conf"
	"github.com/xujintao/gourd/apps/tpl/dao/db"
)

var (
	User  *user
	Users *users
	Repo  *repo
)

func init() {
	// http连接池
	// tp.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //不验证证书
	httpClient := &http.Client{
		// 默认连接池配置
		Transport: http.DefaultTransport.(*http.Transport),

		// 会话超时时间
		Timeout: 10 * time.Second,
	}
	gitlabClient := gitlab.NewClient(httpClient, conf.Config.GetGitlabToken())
	gitlabClient.SetBaseURL(conf.Config.GetGitlabBaseURL())
	db := db.New(conf.Config.GetDBDSN(), conf.Config.GetDBMaxConn())

	// 初始化各业务实例
	User = &user{db, gitlabClient}
	Users = &users{db}
	Repo = &repo{db, gitlabClient}
}
