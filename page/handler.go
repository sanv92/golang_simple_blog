package page

import (
	"net/http"
	"golang_simple_blog/site"
)

func (server *Server) Register(router *site.Router) {
	router.RouterFunc("Home", 1, "/", http.HandlerFunc(server.Home))
	router.RouterFunc("About", 2, "/about", http.HandlerFunc(server.About))
	router.RouterFunc("Contacts", 5, "/contacts", http.HandlerFunc(server.Contacts))
}
