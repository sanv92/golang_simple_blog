package site

import (
	"net/http"
	"sort"
)

type Route struct {
	Title  string
	Weight int
	Path   string
}

type Router struct {
	http.ServeMux
	Routes []Route
}

func (r *Router) RouterFunc(title string, weight int, route string, handler func(w http.ResponseWriter, r *http.Request)) {
	r.HandleFunc(route, handler)
	if weight > 0 {
		r.Routes = append(r.Routes, Route{title, weight, route})

		sort.Slice(r.Routes, func(i, j int) bool {
			return r.Routes[i].Weight < r.Routes[j].Weight
		})
	}
}
