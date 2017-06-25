package news

import (
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"
)


var (
	ErrNotFound = errors.New("not found")
	NoInserterFound = errors.New("No inserter found")
)

type RepoMysql struct {
	DB *sqlx.DB
}

func (repo *RepoMysql) findAll(page, limit int) ([]News, int, error) {
	news := []News{}
	err := repo.DB.Select(&news, "SELECT * FROM news")
	if err != nil {
		return news, 0, ErrNotFound
	}

	length := len(news)

	first := page * limit
	if first < 0 {
		first = 0
	}

	last := first + limit
	if last > length {
		last = length
	}

	if first > length {
		return news, 0, ErrNotFound
	}

	return news[first:last], length, nil
}

func (repo *RepoMysql) findByAlias(alias string) (*News, error) {
	news := News{}
	err := repo.DB.Get(&news, "SELECT * FROM news WHERE alias=?", alias)
	if err != nil {
		return &news, ErrNotFound
	}

	return &news, nil
}

func (repo *RepoMysql) addNew(r *http.Request) bool {
	result := false

	if r.Method  == http.MethodPost {
		r.ParseForm()
		m := &News{
			Title: 		 r.FormValue("title"),
			Alias: 		 r.FormValue("alias"),
			Description: r.FormValue("description"),
			Content: 	 r.FormValue("content"),
		}

		tx, err := repo.DB.Begin()

		_, err = tx.Exec(
			"INSERT INTO news (title, alias, description, content) VALUES (?, ?, ?, ?)",
			m.Title, m.Alias, m.Description, m.Content,
		)

		if err != nil {
			result = false
		}

		tx.Commit()
		result = true
	}

	return result
}

func (repo *RepoMysql) editByAlias(r *http.Request, alias string) {
	method := r.FormValue("_method")

	if method == http.MethodPatch {
		r.ParseForm()
		m := &News{
			Title: 		 r.FormValue("title"),
			Alias: 		 r.FormValue("alias"),
			Description: r.FormValue("description"),
			Content: 	 r.FormValue("content"),
		}

		tx, _ := repo.DB.Begin()

		_, _ = tx.Exec(
			"UPDATE news SET title=?, alias=?, description=?, content=? WHERE alias=? LIMIT 1",
			m.Title, m.Alias, m.Description, m.Content, alias,
		)

		tx.Commit()
	}
}
