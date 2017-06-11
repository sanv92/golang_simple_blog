package json

const NewsFile = "pkg/config/news.json"

type News struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Content      string `json:"content"`
}
