package article

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type ArticleRemoveTaskResult struct {
	app.Result
	Article *Article `json:"article,omitempty"`
}

type ArticleRemoveTask struct {
	app.Task
	Id     int64 `json:"id"`
	Result ArticleRemoveTaskResult
}

func (task *ArticleRemoveTask) GetResult() interface{} {
	return &task.Result
}

func (task *ArticleRemoveTask) GetInhertType() string {
	return "article"
}

func (task *ArticleRemoveTask) GetClientName() string {
	return "Article.Remove"
}
