package news

import (
	"net/http"
)

func (server *Server) Register(mux *http.ServeMux) {
	mux.HandleFunc("/news/", server.List)
	mux.HandleFunc("/news/show/", server.Full)
}
