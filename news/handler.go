package news

import (
	"net/http"
	"github.com/SanderV1992/golang_simple_blog/site"
)

func (server *Server) Register(router *site.Router) {
	router.RouterFunc("News", 3, "/news/", http.HandlerFunc(server.List))
	router.RouterFunc("NewsFull", -1, "/news/show/", http.HandlerFunc(server.Full))
}
