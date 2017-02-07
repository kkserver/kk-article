package article

import (
	"database/sql"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/app/remote"
)

type Article struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Author  string `json:"author"`
	Origin  string `json:"origin"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Image   string `json:"image"`
	Images  string `json:"images"`
	Type    string `json:"type"`
	Content string `json:"content,omitempty"`
	Tags    string `json:"tags"`
	InTime  int64  `json:"intime"`
	Mtime   int64  `json:"mtime"`
	Ctime   int64  `json:"ctime"`
}

type IArticleApp interface {
	app.IApp
	GetDB() (*sql.DB, error)
	GetPrefix() string
	GetArticleTable() *kk.DBTable
}

type ArticleApp struct {
	app.App
	DB *app.DBConfig

	Remote *remote.Service

	Article      *ArticleService
	ArticleTable kk.DBTable
}

func (C *ArticleApp) GetDB() (*sql.DB, error) {
	return C.DB.Get(C)
}

func (C *ArticleApp) GetPrefix() string {
	return C.DB.Prefix
}

func (C *ArticleApp) GetArticleTable() *kk.DBTable {
	return &C.ArticleTable
}
