package article

import (
	"bytes"
	"fmt"
	"github.com/kkserver/kk-lib/kk"
	"github.com/kkserver/kk-lib/kk/app"
	"github.com/kkserver/kk-lib/kk/dynamic"
	"time"
)

type ArticleService struct {
	app.Service

	Get      *ArticleTask
	Set      *ArticleSetTask
	Remove   *ArticleRemoveTask
	Create   *ArticleCreateTask
	Query    *ArticleQueryTask
	Exchange *ArticleExchangeTask
}

func (S *ArticleService) Handle(a app.IApp, task app.ITask) error {
	return app.ServiceReflectHandle(a, task, S)
}

func (S *ArticleService) HandleArticleCreateTask(a IArticleApp, task *ArticleCreateTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_ARTICLE
		task.Result.Errmsg = err.Error()
		return nil
	}

	if task.Name != "" {

		count, err := kk.DBQueryCount(db, a.GetArticleTable(), a.GetPrefix(), " WHERE name=?", task.Name)

		if err != nil {
			task.Result.Errno = ERROR_ARTICLE
			task.Result.Errmsg = err.Error()
			return nil
		}

		if count > 0 {
			task.Result.Errno = ERROR_ARTICLE_NAME
			task.Result.Errmsg = "Article name already exists"
			return nil
		}
	}

	v := Article{}

	v.Uid = task.Uid
	v.Alias = task.Alias
	v.Title = task.Title
	v.Image = task.Image
	v.Images = task.Images
	v.Name = task.Name
	v.Author = task.Author
	v.Type = task.Type
	v.Content = task.Content
	v.Origin = task.Origin
	v.Tags = task.Tags
	v.Summary = task.Summary
	v.Ctime = time.Now().Unix()
	v.Mtime = v.Ctime
	v.Oid = NewOid()

	_, err = kk.DBInsert(db, a.GetArticleTable(), a.GetPrefix(), &v)

	if err != nil {
		task.Result.Errno = ERROR_ARTICLE
		task.Result.Errmsg = err.Error()
		return nil
	}

	task.Result.Article = &v

	return nil
}

func (S *ArticleService) HandleArticleSetTask(a IArticleApp, task *ArticleSetTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_ARTICLE
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Article{}

	if task.Id != 0 {

		rows, err := kk.DBQuery(db, a.GetArticleTable(), a.GetPrefix(), " WHERE id=?", task.Id)

		if err != nil {
			task.Result.Errno = ERROR_ARTICLE
			task.Result.Errmsg = err.Error()
			return nil
		}

		defer rows.Close()

		if rows.Next() {
			scanner := kk.NewDBScaner(&v)
			err = scanner.Scan(rows)
			if err != nil {
				task.Result.Errno = ERROR_ARTICLE
				task.Result.Errmsg = err.Error()
				return nil
			}
		} else {
			task.Result.Errno = ERROR_ARTICLE_NOT_FOUND
			task.Result.Errmsg = "Not Found Article"
			return nil
		}

	} else if task.Name != "" {

		rows, err := kk.DBQuery(db, a.GetArticleTable(), a.GetPrefix(), " WHERE name=?", task.Name)

		if err != nil {
			task.Result.Errno = ERROR_ARTICLE
			task.Result.Errmsg = err.Error()
			return nil
		}

		defer rows.Close()

		if rows.Next() {
			scanner := kk.NewDBScaner(&v)
			err = scanner.Scan(rows)
			if err != nil {
				task.Result.Errno = ERROR_ARTICLE
				task.Result.Errmsg = err.Error()
				return nil
			}
		} else {
			v.Name = task.Name
		}

	}

	keys := map[string]bool{}

	if task.Title != nil {
		v.Title = dynamic.StringValue(task.Title, v.Title)
		keys["title"] = true
	}

	if task.Image != nil {
		v.Image = dynamic.StringValue(task.Image, v.Image)
		keys["image"] = true
	}

	if task.Images != nil {
		v.Images = dynamic.StringValue(task.Images, v.Images)
		keys["images"] = true
	}

	if task.Author != nil {
		v.Author = dynamic.StringValue(task.Author, v.Author)
		keys["author"] = true
	}

	if task.Type != nil {
		v.Type = dynamic.StringValue(task.Type, v.Type)
		keys["type"] = true
	}

	if task.Origin != nil {
		v.Origin = dynamic.StringValue(task.Origin, v.Origin)
		keys["origin"] = true
	}

	if task.Tags != nil {
		v.Tags = dynamic.StringValue(task.Tags, v.Tags)
		keys["tags"] = true
	}

	if task.Content != nil {
		v.Content = dynamic.StringValue(task.Content, v.Content)
		keys["content"] = true
	}

	if task.Summary != nil {
		v.Summary = dynamic.StringValue(task.Summary, v.Summary)
		keys["summary"] = true
	}

	if task.Uid != nil {
		v.Uid = dynamic.IntValue(task.Uid, v.Uid)
		keys["uid"] = true
	}

	v.Mtime = time.Now().Unix()

	keys["mtime"] = true

	if v.Id == 0 {
		_, err = kk.DBInsert(db, a.GetArticleTable(), a.GetPrefix(), &v)

		if err != nil {
			task.Result.Errno = ERROR_ARTICLE
			task.Result.Errmsg = err.Error()
			return nil
		}
	} else {

		_, err = kk.DBUpdateWithKeys(db, a.GetArticleTable(), a.GetPrefix(), &v, keys)

		if err != nil {
			task.Result.Errno = ERROR_ARTICLE
			task.Result.Errmsg = err.Error()
			return nil
		}

	}

	task.Result.Article = &v

	return nil
}

func (S *ArticleService) HandleArticleTask(a IArticleApp, task *ArticleTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_ARTICLE
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Article{}

	if task.Id != 0 {

		rows, err := kk.DBQuery(db, a.GetArticleTable(), a.GetPrefix(), " WHERE id=?", task.Id)

		if err != nil {
			task.Result.Errno = ERROR_ARTICLE
			task.Result.Errmsg = err.Error()
			return nil
		}

		defer rows.Close()

		if rows.Next() {
			scanner := kk.NewDBScaner(&v)
			err = scanner.Scan(rows)
			if err != nil {
				task.Result.Errno = ERROR_ARTICLE
				task.Result.Errmsg = err.Error()
				return nil
			}
		} else {
			task.Result.Errno = ERROR_ARTICLE_NOT_FOUND
			task.Result.Errmsg = "Not Found Article"
			return nil
		}

	} else if task.Name != "" {

		rows, err := kk.DBQuery(db, a.GetArticleTable(), a.GetPrefix(), " WHERE name=?", task.Name)

		if err != nil {
			task.Result.Errno = ERROR_ARTICLE
			task.Result.Errmsg = err.Error()
			return nil
		}

		defer rows.Close()

		if rows.Next() {
			scanner := kk.NewDBScaner(&v)
			err = scanner.Scan(rows)
			if err != nil {
				task.Result.Errno = ERROR_ARTICLE
				task.Result.Errmsg = err.Error()
				return nil
			}
		} else {
			task.Result.Errno = ERROR_ARTICLE_NOT_FOUND
			task.Result.Errmsg = "Not Found Article"
			return nil
		}

	}

	task.Result.Article = &v

	return nil
}

func (S *ArticleService) HandleArticleRemoveTask(a IArticleApp, task *ArticleRemoveTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_ARTICLE
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Article{}

	rows, err := kk.DBQuery(db, a.GetArticleTable(), a.GetPrefix(), " WHERE id=?", task.Id)

	if err != nil {
		task.Result.Errno = ERROR_ARTICLE
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	if rows.Next() {
		scanner := kk.NewDBScaner(&v)
		err = scanner.Scan(rows)
		if err != nil {
			task.Result.Errno = ERROR_ARTICLE
			task.Result.Errmsg = err.Error()
			return nil
		}

		_, err = kk.DBDelete(db, a.GetArticleTable(), a.GetPrefix(), " WHERE id=?", task.Id)

		if err != nil {
			task.Result.Errno = ERROR_ARTICLE
			task.Result.Errmsg = err.Error()
			return nil
		}
	} else {
		task.Result.Errno = ERROR_ARTICLE_NOT_FOUND
		task.Result.Errmsg = "Not Found Article"
		return nil
	}

	task.Result.Article = &v

	return nil
}

func (S *ArticleService) HandleArticleQueryTask(a IArticleApp, task *ArticleQueryTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_ARTICLE
		task.Result.Errmsg = err.Error()
		return nil
	}

	var articles = []Article{}

	var args = []interface{}{}

	var sql = bytes.NewBuffer(nil)

	sql.WriteString(" WHERE 1")

	if task.Id != 0 {
		sql.WriteString(" AND id=?")
		args = append(args, task.Id)
	}

	if task.Uid != 0 {
		sql.WriteString(" AND uid=?")
		args = append(args, task.Uid)
	}

	if task.Alias != "" {
		sql.WriteString(" AND alias=?")
		args = append(args, task.Alias)
	}

	if task.Name != "" {
		sql.WriteString(" AND name=?")
		args = append(args, task.Name)
	}

	if task.Keyword != "" {
		q := "%" + task.Keyword + "%"
		sql.WriteString(" AND (title LIKE ? OR author LIKE ? OR summary LIKE ?)")
		args = append(args, q, q, q)
	}

	if task.OrderBy == "asc" {
		sql.WriteString(" ORDER BY id ASC")
	} else if task.OrderBy == "oid" {
		sql.WriteString(" ORDER BY oid DESC,id DESC")
	} else {
		sql.WriteString(" ORDER BY id DESC")
	}

	var pageIndex = task.PageIndex
	var pageSize = task.PageSize

	if pageIndex < 1 {
		pageIndex = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	if task.Counter {
		var counter = ArticleQueryCounter{}
		counter.PageIndex = pageIndex
		counter.PageSize = pageSize
		counter.RowCount, err = kk.DBQueryCount(db, a.GetArticleTable(), a.GetPrefix(), sql.String(), args...)
		if err != nil {
			task.Result.Errno = ERROR_ARTICLE
			task.Result.Errmsg = err.Error()
			return nil
		}
		if counter.RowCount%pageSize == 0 {
			counter.PageCount = counter.RowCount / pageSize
		} else {
			counter.PageCount = counter.RowCount/pageSize + 1
		}
		task.Result.Counter = &counter
	}

	sql.WriteString(fmt.Sprintf(" LIMIT %d,%d", (pageIndex-1)*pageSize, pageSize))

	var v = Article{}
	var scanner = kk.NewDBScaner(&v)

	rows, err := db.Query(fmt.Sprintf("SELECT id,name,author,origin,title,image,images,type,intime,mtime,ctime FROM %s%s %s", a.GetPrefix(), a.GetArticleTable().Name, sql.String()), args...)

	if err != nil {
		task.Result.Errno = ERROR_ARTICLE
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	for rows.Next() {

		err = scanner.Scan(rows)

		if err != nil {
			task.Result.Errno = ERROR_ARTICLE
			task.Result.Errmsg = err.Error()
			return nil
		}

		articles = append(articles, v)
	}

	task.Result.Articles = articles

	return nil
}

func (S *ArticleService) HandleArticleExchangeTask(a IArticleApp, task *ArticleExchangeTask) error {

	var db, err = a.GetDB()

	if err != nil {
		task.Result.Errno = ERROR_ARTICLE
		task.Result.Errmsg = err.Error()
		return nil
	}

	v := Article{}
	vs := []Article{}

	rows, err := db.Query(fmt.Sprintf("SELECT id,oid FROM %s%s WHERE id IN (?,?)", a.GetPrefix(), a.GetArticleTable().Name), task.FromId, task.ToId)

	if err != nil {
		task.Result.Errno = ERROR_ARTICLE
		task.Result.Errmsg = err.Error()
		return nil
	}

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&v.Id, &v.Oid)

		if err != nil {
			task.Result.Errno = ERROR_ARTICLE
			task.Result.Errmsg = err.Error()
			return nil
		}

		vs = append(vs, v)
	}

	if len(vs) > 1 {

		_, err = db.Exec(fmt.Sprintf("UPDATE %s%s SET oid=? WHERE id=?", a.GetPrefix(), a.GetArticleTable().Name), vs[0].Oid, vs[1].Id)

		if err != nil {
			task.Result.Errno = ERROR_ARTICLE
			task.Result.Errmsg = err.Error()
			return nil
		}

		_, err = db.Exec(fmt.Sprintf("UPDATE %s%s SET oid=? WHERE id=?", a.GetPrefix(), a.GetArticleTable().Name), vs[1].Oid, vs[0].Id)

		if err != nil {
			task.Result.Errno = ERROR_ARTICLE
			task.Result.Errmsg = err.Error()
			return nil
		}

	}

	return nil
}
