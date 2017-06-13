package news

import (
	"net/http"
	"strconv"

	"blog2/site"
)


type News struct {
	Title        string `json:"title"`
	Alias        string `json:"alias"`
	Description  string `json:"description"`
	Content      string `json:"content"`
}

type Server struct{
	*site.Renderer
}

func (server *Server) List(w http.ResponseWriter, r *http.Request) {
	queryPage  := r.URL.Query().Get("page")
	if queryPage == "" {
		queryPage = "1"
	}
	pageNum, _ := strconv.Atoi(queryPage)
	PageLimit  := 3

	var news []News
	site.ReadJSON("config/news.json", &news)

	startPoint := ((pageNum * PageLimit) - 3)
	if startPoint <= 0 {
		startPoint = 0
	}

	endPoint   := pageNum * PageLimit
	news_slice := news[startPoint:endPoint]

	p := site.NewPagination(len(news), PageLimit, pageNum)

	data := struct {
		News    []News
		Pagination site.Pagination
	}{
		news_slice,
		*p,
	}

	server.Render(w, "news_list.tmpl", data)
}

func (server *Server) Full(w http.ResponseWriter, r *http.Request) {
	newsName := r.URL.Query().Get("id")

	var news []News
	site.ReadJSON("config/news.json", &news)

	found := false
	var foundNews News
	for _, item := range news {
		if item.Alias == newsName {
			foundNews = item
			found = true
			break
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		foundNews = News{Title: "Not Found 404"}
	}

	data := struct {
		News    News
	}{
		foundNews,
	}
	server.Render(w, "news_full.tmpl", data)
}
