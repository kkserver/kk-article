package article

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type ArticleQueryCounter struct {
	PageIndex int `json:"p"`
	PageSize  int `json:"size"`
	PageCount int `json:"count"`
	RowCount  int `json:"rowCount"`
}

type ArticleQueryTaskResult struct {
	app.Result
	Counter  *ArticleQueryCounter `json:"counter,omitempty"`
	Articles []Article            `json:"articles,omitempty"`
}

type ArticleQueryTask struct {
	app.Task
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Alias     string `json:"alias"`
	Uid       int64  `json:"uid"`
	Keyword   string `json:"q"`
	OrderBy   string `json:"orderBy"` // desc, asc
	PageIndex int    `json:"p"`
	PageSize  int    `json:"size"`
	Counter   bool   `json:"counter"`
	Result    ArticleQueryTaskResult
}

func (task *ArticleQueryTask) GetResult() interface{} {
	return &task.Result
}

func (task *ArticleQueryTask) GetInhertType() string {
	return "article"
}

func (task *ArticleQueryTask) GetClientName() string {
	return "Article.Query"
}
