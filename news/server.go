package news

import (
	"net/http"

	"github.com/SanderV1992/golang_simple_blog/site"
	"strconv"
)


type Repo interface {
	findAll(page, limit int) ([]News, int, error)
	findByAlias(alias string) (*News, error)
	addNew(r *http.Request) (bool)
	editByAlias(r *http.Request, alias string)
}

type Server struct {
	*site.Renderer
	Repo Repo
}

func (server *Server) List(w http.ResponseWriter, r *http.Request) {
	queryPage  := r.URL.Query().Get("page")

	page, err := strconv.Atoi(queryPage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		page = 1
	}
	limit := 3

	news, count, err := server.Repo.findAll(page - 1, limit)
	if err != nil {
		site.Abort(404, w, r)
		return
	}

	p := site.NewPagination(count, limit, page)

	data := struct {
		News    []News
		Pagination site.Pagination
	}{
		news,
		*p,
	}

	server.Render(w, "news_list.tmpl", data)
}

func (server *Server) Full(w http.ResponseWriter, r *http.Request) {
	alias := r.URL.Query().Get("id")
	news, err := server.Repo.findByAlias(alias)
	if err != nil {
		site.Abort(404, w, r)
		return
	}

	data := struct {
		News    News
	}{
		*news,
	}
	server.Render(w, "news_full.tmpl", data)
}

func (server *Server) Add(w http.ResponseWriter, r *http.Request) {
	result := server.Repo.addNew(r)
	server.Render(w, "news_add.tmpl", result)
}

func (server *Server) Edit(w http.ResponseWriter, r *http.Request) {
	alias := r.URL.Query().Get("id")
	server.Repo.editByAlias(r, alias)

	news, err := server.Repo.findByAlias(alias)
	if err != nil {
		site.Abort(404, w, r)
		return
	}

	data := struct {
		News    News
	}{
		*news,
	}

	server.Render(w, "news_edit.tmpl", data)
}
