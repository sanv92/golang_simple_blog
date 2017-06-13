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

// PageRenderer ////////////////////////
//
type PageRenderer struct{
	path      string
	templates *template.Template
}

func NewRenderer(templatespath string) (*PageRenderer, error) {
	renderer := &PageRenderer{}
	renderer.path = templatespath
	return renderer, renderer.loadTemplates()
}

func (renderer *PageRenderer) loadTemplates() error {
	var err error
	renderer.templates, err = template.New("test").Funcs(renderer.funcs()).ParseGlob(renderer.path)
	return err
}

func (renderer *PageRenderer) funcs() template.FuncMap {
	return template.FuncMap{
		"loop": func(n int) []struct{} {
			return make([]struct{}, n)
		},
		"add": func(x, y int) int {
			return x + y
		},
		"Menu": func() []Menu {
			var menu []Menu
			ReadJSON("config/menu.json", &menu)
			return menu
		},
	}
}

func (renderer *PageRenderer) Render(w io.Writer, name string, data interface{}) error {
	err := renderer.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Println(err)
	}
	return err
}
