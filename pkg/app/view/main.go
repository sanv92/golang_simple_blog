package view

import (
	//"fmt"

	"net/http"
	"text/template"
	"errors"

)

var (
	ErrTemplateDoesNotExist = errors.New("The template does not exist.")
)

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templatesDir := "pkg/templates/"

	tplHome     := []string{templatesDir + "home.tmpl", templatesDir + "base.tmpl"}
	tplNewsList := []string{templatesDir + "news_list.tmpl", templatesDir + "base.tmpl"}
	tplNewsFull := []string{templatesDir + "news_full.tmpl", templatesDir + "base.tmpl"}
	tplAbout    := []string{templatesDir + "about.tmpl", templatesDir + "base.tmpl"}
	tplContacts := []string{templatesDir + "contacts.tmpl", templatesDir + "base.tmpl"}

	tplHome     = append(tplHome, templatesDir + "header.tmpl", templatesDir + "footer.tmpl")
	tplNewsList = append(tplNewsList, templatesDir + "header.tmpl", templatesDir + "footer.tmpl")
	tplNewsFull = append(tplNewsFull, templatesDir + "header.tmpl", templatesDir + "footer.tmpl")
	tplAbout    = append(tplAbout, templatesDir + "header.tmpl", templatesDir + "footer.tmpl")
	tplContacts = append(tplContacts, templatesDir + "header.tmpl", templatesDir + "footer.tmpl")

	templates["home"]      = template.Must(template.ParseFiles(tplHome...))
	templates["news_list"] = template.Must(template.ParseFiles(tplNewsList...))
	templates["news_full"] = template.Must(template.ParseFiles(tplNewsFull...))
	templates["about"]     = template.Must(template.ParseFiles(tplAbout...))
	templates["contacts"]  = template.Must(template.ParseFiles(tplContacts...))

	templates["header"]      = template.Must(template.ParseFiles(templatesDir + "header.tmpl"))
}

type View string
type Data struct {
	Menu    interface{}
	Content interface{}
}

func (v View) Render(data interface{}, w http.ResponseWriter, r *http.Request) error {
	tpl, ok := templates[string(v)]
	if !ok {
		return ErrTemplateDoesNotExist
	}

	//tpl.Execute(w, &Data{Menu: "aaa", Content: "bbb"})
	tpl.ExecuteTemplate(w, "base", data)

	return nil
}
