package article

import (
	"github.com/kkserver/kk-lib/kk/app"
)

type ArticleExchangeTaskResult struct {
	app.Result
}

type ArticleExchangeTask struct {
	app.Task
	FromId int64 `json:"fromId"`
	ToId   int64 `json:"toId"` // 0 时至顶
	Result ArticleExchangeTaskResult
}

func (task *ArticleExchangeTask) GetResult() interface{} {
	return &task.Result
}

func (task *ArticleExchangeTask) GetInhertType() string {
	return "article"
}

func (task *ArticleExchangeTask) GetClientName() string {
	return "Article.Exchange"
}
