package news

import (
	"net/http"

	"strconv"

	"github.com/SanderV1992/golang_simple_blog/site"
)

type Repo interface {
	FindAll(page, limit int) ([]News, int, error)
	FindByAlias(alias string) (*News, error)
	create(data News) error
	update(data News, alias string)
}

type Server struct {
	*site.Renderer
	Repo Repo
}

func (server *Server) List(w http.ResponseWriter, r *http.Request) {
	queryPage := r.URL.Query().Get("page")

	page, err := strconv.Atoi(queryPage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		page = 1
	}
	limit := 3

	news, count, err := server.Repo.FindAll(page-1, limit)
	if err != nil {
		site.Abort(404, w, r)
		return
	}

	p := site.NewPagination(count, limit, page)

	data := struct {
		News       []News
		Pagination site.Pagination
	}{
		news,
		*p,
	}

	server.Render(w, "news_list.tmpl", data)
}

func (server *Server) Full(w http.ResponseWriter, r *http.Request) {
	alias := r.URL.Query().Get("id")
	news, err := server.Repo.FindByAlias(alias)
	if err != nil {
		site.Abort(404, w, r)
		return
	}

	data := struct {
		News News
	}{
		*news,
	}
	server.Render(w, "news_full.tmpl", data)
}

func (server *Server) Add(w http.ResponseWriter, r *http.Request) {

	var result bool
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			site.Abort(404, w, r)
			return
		}

		values := &News{
			Title:       r.FormValue("title"),
			Alias:       r.FormValue("alias"),
			Description: r.FormValue("description"),
			Content:     r.FormValue("content"),
		}

		if err := server.Repo.create(*values); err != nil {
			site.Abort(404, w, r)
		}
	}

	server.Render(w, "news_add.tmpl", result)
}

func (server *Server) Edit(w http.ResponseWriter, r *http.Request) {
	alias := r.URL.Query().Get("id")

	// update
	method := r.FormValue("_method")
	if method == http.MethodPatch {
		if err := r.ParseForm(); err != nil {
			site.Abort(404, w, r)
			return
		}

		values := &News{
			Title:       r.FormValue("title"),
			Alias:       r.FormValue("alias"),
			Description: r.FormValue("description"),
			Content:     r.FormValue("content"),
		}

		server.Repo.update(*values, alias)
	}

	// select
	news, err := server.Repo.FindByAlias(alias)
	if err != nil {
		site.Abort(404, w, r)
		return
	}

	data := struct {
		News News
	}{
		*news,
	}

	server.Render(w, "news_edit.tmpl", data)
}
