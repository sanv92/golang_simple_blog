package page

import (
	"net/http"
)

func (server *Server) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", server.Home)
	mux.HandleFunc("/about", server.About)
	mux.HandleFunc("/contacts", server.Contacts)
}
