package article

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type ArticleTaskResult struct {
	app.Result
	Article *Article `json:"article,omitempty"`
}

type ArticleTask struct {
	app.Task
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Result ArticleSetTaskResult
}

func (task *ArticleTask) GetResult() interface{} {
	return &task.Result
}

func (task *ArticleTask) GetInhertType() string {
	return "article"
}

func (task *ArticleTask) GetClientName() string {
	return "Article.Get"
}
