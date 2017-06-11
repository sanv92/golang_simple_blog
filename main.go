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


func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/contacts", ContactsPage)
	http.HandleFunc("/news/", NewsList)
	http.HandleFunc("/news/show/", NewsFull)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	fmt.Println("Listen on port: 8080")
	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	//var menu []json.Menu
	//ReadJSON(&menu)

	tpl.ExecuteTemplate(w, "home.tmpl", "11122")
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.tmpl", nil)
}

func ContactsPage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "contacts.tmpl", nil)
}

func NewsList(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "news_list.tmpl", "2222")
}

func NewsFull(w http.ResponseWriter, r *http.Request) {
	//newsID := strings.SplitN(req.URL.Path, "/", 4)[3]
	//fmt.Fprint(w, string(newsID))

	tpl.ExecuteTemplate(w, "news_full.tmpl", "44444")
}

type Menu struct {
	Name   string `json:"name"`
	Alias  string `json:"alias"`
}

type News struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Content      string `json:"content"`
}

func ReadJSON(file string, result interface{}) error {
	file, err := os.Open(file)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(result)
}
