package news

import (
	"net/http"
	"golang_simple_blog/site"
)

func (server *Server) Register(mux *http.ServeMux, router *site.Router) {
	//router.RouterFunc("About", 222, "/about")

	mux.HandleFunc("/news/", server.List)
	mux.HandleFunc("/news/show/", server.Full)
}
