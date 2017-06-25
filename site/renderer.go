package site

import (
	"text/template"
	"log"
	"io"
	"errors"
)


var (
	ErrTemplateDoesNotExist = errors.New("The template does not exist.")
)


type Renderer struct {
	path      string
	templates *template.Template
}

func NewRenderer(templatespath string, router *Router) (*Renderer, error) {
	renderer     := &Renderer{}
	renderer.path = templatespath

	return renderer, renderer.loadTemplates(
		router,
	)
}

func (renderer *Renderer) loadTemplates(router *Router) error {
	var err error
	renderer.templates, err = template.New("test").Funcs(renderer.funcs(
		router,
	)).ParseGlob(renderer.path)
	return err
}

func (renderer *Renderer) funcs(router *Router) template.FuncMap {

	return template.FuncMap{
		"loop": func(n int) []struct{} {
			return make([]struct{}, n)
		},
		"add": func(x, y int) int {
			return x + y
		},
		"Menu": func() []Route {
			return router.Routes
		},
	}
}

func (renderer *Renderer) Render(w io.Writer, name string, data interface{}) error {
	err := renderer.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Println(err)
	}
	return err
}
