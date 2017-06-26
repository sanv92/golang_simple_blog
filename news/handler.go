package news

import (
	"net/http"

	"github.com/SanderV1992/golang_simple_blog/site"
)

func (server *Server) Register(router *site.Router) {
	router.RouterFunc("News", 3, "/news/", http.HandlerFunc(server.List))
	router.RouterFunc("NewsFull", -1, "/news/show/", http.HandlerFunc(server.Full))

	router.RouterFunc("News add", 4, "/news/add/", http.HandlerFunc(server.Add))
	router.RouterFunc("News edit", -1, "/news/edit/", http.HandlerFunc(server.Edit))
}
