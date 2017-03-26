package article

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type ArticleCreateTaskResult struct {
	app.Result
	Article *Article `json:"article,omitempty"`
}

type ArticleCreateTask struct {
	app.Task
	Uid     int64  `json:"uid"`
	Name    string `json:"name"`
	Alias   string `json:"alias"`
	Author  string `json:"author"`
	Origin  string `json:"origin"`
	Image   string `json:"image"`
	Images  string `json:"images"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Type    string `json:"type"`
	Content string `json:"content"`
	Tags    string `json:"tags"`
	Result  ArticleCreateTaskResult
}

func (task *ArticleCreateTask) GetResult() interface{} {
	return &task.Result
}

func (task *ArticleCreateTask) GetInhertType() string {
	return "article"
}

func (task *ArticleCreateTask) GetClientName() string {
	return "Article.Create"
}
