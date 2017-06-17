package page

import (
	"net/http"
	"golang_simple_blog/site"
)

func (server *Server) Register(mux *http.ServeMux, router *site.Router) {
	router.RouterFunc("Home", 10, "/", http.HandlerFunc(server.Home))
	router.RouterFunc("About", 50, "/about", http.HandlerFunc(server.About))
	router.RouterFunc("Contacts", 4, "/contacts", http.HandlerFunc(server.Contacts))

	/*
	mux.HandleFunc("/", server.Home)
	mux.HandleFunc("/about", server.About)
	mux.HandleFunc("/contacts", server.Contacts)
	*/
}
