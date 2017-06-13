package page

import (
	"net/http"

	. "blog2/app/site"
)


type Server struct{
	*PageRenderer
}

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	server.Render(w, "home.tmpl", nil)
}

func (server *Server) About(w http.ResponseWriter, r *http.Request) {
	server.Render(w, "about.tmpl", nil)
}

func (server *Server) Contacts(w http.ResponseWriter, r *http.Request) {
	server.Render(w, "contacts.tmpl", nil)
}
