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


type Renderer struct{
	path      string
	templates *template.Template
}

func NewRenderer(templatespath string) (*Renderer, error) {
	renderer := &Renderer{}
	renderer.path = templatespath
	return renderer, renderer.loadTemplates()
}

func (renderer *Renderer) loadTemplates() error {
	var err error
	renderer.templates, err = template.New("test").Funcs(renderer.funcs()).ParseGlob(renderer.path)
	return err
}

func (renderer *Renderer) funcs() template.FuncMap {
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

func (renderer *Renderer) Render(w io.Writer, name string, data interface{}) error {
	err := renderer.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Println(err)
	}
	return err
}
