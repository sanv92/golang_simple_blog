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

var tpl = template.Must(template.ParseGlob("templates/*"))


type View string

func (v View) Render(data interface{}, w http.ResponseWriter, r *http.Request) error {
	if err := tpl.ExecuteTemplate(w, string(v) + ".tmpl", data); err != nil {
		return ErrTemplateDoesNotExist
	}

	return nil
}


func HomePage(w http.ResponseWriter, r *http.Request) {
	//var menu []json.Menu
	//ReadJSON(&menu)

	View("home").Render("1111", w, r)
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	View("about").Render(nil, w, r)
}

func ContactsPage(w http.ResponseWriter, r *http.Request) {
	View("contacts").Render(nil, w, r)
}

func NewsList(w http.ResponseWriter, r *http.Request) {
	View("news_list").Render("2222", w, r)
}

func NewsFull(w http.ResponseWriter, r *http.Request) {
	//newsID := strings.SplitN(req.URL.Path, "/", 4)[3]
	//fmt.Fprint(w, string(newsID))

	View("news_full").Render("4444", w, r)
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
