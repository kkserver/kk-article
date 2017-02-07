package article

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type ArticleSetTaskResult struct {
	app.Result
	Article *Article `json:"article,omitempty"`
}

type ArticleSetTask struct {
	app.Task
	Id   int64  `json:"id"`
	Name string `json:"name"`

	Author  interface{} `json:"author"`
	Origin  interface{} `json:"origin"`
	Title   interface{} `json:"title"`
	Summary interface{} `json:"summary"`
	Image   interface{} `json:"image"`
	Images  interface{} `json:"images"`
	Type    interface{} `json:"type"`
	Content interface{} `json:"content"`
	Tags    interface{} `json:"tags"`

	Result ArticleSetTaskResult
}

func (task *ArticleSetTask) GetResult() interface{} {
	return &task.Result
}

func (task *ArticleSetTask) GetInhertType() string {
	return "article"
}

func (task *ArticleSetTask) GetClientName() string {
	return "Article.Set"
}
