package news

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type RepoMysql struct {
	DB *sqlx.DB
}

func (repo *RepoMysql) findAll(page, limit int) ([]News, int, error) {
	news := []News{}
	err := repo.DB.Select(&news, "SELECT * FROM news")
	if err != nil {
		fmt.Errorf("DB select news fail: %v", err)
	}

	first := page * limit
	if first < 0 {
		first = 0
	}

	last := first + limit
	if last > len(news) {
		last = len(news)
	}

	return news[first:last], len(news), nil
}

func (repo *RepoMysql) findByAlias(alias string) (News, error) {
	news := News{}
	err := repo.DB.Get(&news, "SELECT * FROM news WHERE alias=?", alias)
	if err != nil {
		fmt.Errorf("DB select news fail: %v", err)
	}

	return news, nil
}
