package page

import (
	"net/http"

	. "blog2/app/site"
)

// Page Server ////////////////////////
//
type PageServer struct{
	*PageRenderer
}

func (server *PageServer) Home(w http.ResponseWriter, r *http.Request) {
	server.Render(w, "home.tmpl", nil)
}

func (server *PageServer) About(w http.ResponseWriter, r *http.Request) {
	server.Render(w, "about.tmpl", nil)
}

func (server *PageServer) Contacts(w http.ResponseWriter, r *http.Request) {
	server.Render(w, "contacts.tmpl", nil)
}
