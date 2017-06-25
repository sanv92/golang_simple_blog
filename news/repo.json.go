package news

import (
	"errors"

	"github.com/SanderV1992/golang_simple_blog/site"
)


type RepoJson struct {}

func (repo *RepoJson) findAll(page, limit int) ([]News, int, error) {
	// ning siin kasutada repo.DB.Exec
	var news []News
	site.ReadJSON("config/news.json", &news)

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

func (repo *RepoJson) findByAlias(alias string) (*News, error) {
	var news []News
	site.ReadJSON("config/news.json", &news)

	for _, item := range news {
		if item.Alias == alias {
			return &item, nil
		}
	}

	return nil, errors.New("not found")
}


func (repo *RepoJson) create(data News) bool { return false }
func (repo *RepoJson) update(data News, alias string) {}
