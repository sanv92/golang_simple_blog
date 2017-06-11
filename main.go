package main

import (
	"fmt"
	"os"
	"net/http"
	"text/template"
	"encoding/json"
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

	templatesDir := "templates/"

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

	//templates["header"]      = template.Must(template.ParseFiles(templatesDir + "header.tmpl"))
}

type View string

func (v View) Render(data interface{}, w http.ResponseWriter, r *http.Request) error {
	tpl, ok := templates[string(v)]
	if !ok {
		return ErrTemplateDoesNotExist
	}

	tpl.ExecuteTemplate(w, "base", data)

	return nil
}


func HomePage(w http.ResponseWriter, r *http.Request) {
	//var menu []json.Menu
	//ReadJSON(&menu)

	View("home").Render(nil, w, r)
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	View("about").Render(nil, w, r)
}

func ContactsPage(w http.ResponseWriter, r *http.Request) {
	View("contacts").Render(nil, w, r)
}

func NewsList(w http.ResponseWriter, r *http.Request) {
	View("news_list").Render("lalal", w, r)
}

func NewsFull(w http.ResponseWriter, r *http.Request) {
	//newsID := strings.SplitN(req.URL.Path, "/", 4)[3]
	//fmt.Fprint(w, string(newsID))

	View("news_full").Render("lalal", w, r)
}


const (
	NewsFile = "pkg/config/news.json"
	MenuFile = "pkg/config/menu.json"
)

type Menu struct {
	Name   string `json:"name"`
	Alias  string `json:"alias"`
}

type News struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Content      string `json:"content"`
}

func ReadJSON(result interface{}) error {
	file, err := os.Open(MenuFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(result)
}


func middlewareMenu(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//var menu []json.Menu
		//ReadJSON(&menu)

		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", middlewareMenu(HomePage))
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/contacts", ContactsPage)
	http.HandleFunc("/news/", NewsList)
	http.HandleFunc("/news/show/", NewsFull)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	fmt.Println("Listen on port: 8080")
	http.ListenAndServe(":8080", nil)
}
