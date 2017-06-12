package main

import (
	"fmt"
	"os"
	"net/http"
	"text/template"
	"encoding/json"
	"errors"
	"strconv"
	"log"
	"io"
)

func main() {
	renderer, err := NewRenderer("templates/*")
	if err != nil {
		log.Fatal(err)
	}
	pages := &PageServer{renderer}
	news := &NewsServer{renderer}

	http.HandleFunc("/", pages.Home)
	http.HandleFunc("/about", pages.About)
	http.HandleFunc("/contacts", pages.Contacts)

	http.HandleFunc("/news/", news.List)
	http.HandleFunc("/news/show/", news.Full)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	fmt.Println("Listen on port: 8080")
	http.ListenAndServe(":8080", nil)
}

var (
	ErrTemplateDoesNotExist = errors.New("The template does not exist.")
)

func NewRenderer(templatespath string) (*PageRenderer, error) {
	renderer := &PageRenderer{}
	renderer.path = templatespath
	return renderer, renderer.loadTemplates()
}

// PageRenderer ////////////////////////
//
type PageRenderer struct{
	path      string
	templates *template.Template
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


// Page Server ////////////////////////
//
type PageServer struct{
	*PageRenderer
}

func (server *PageServer) Home(w http.ResponseWriter, r *http.Request) {
	server.Render(w, "home.tmpl", nil)
}

func (server *PageServer) About(w http.ResponseWriter, r *http.Request) {
	server.Render(w, "about.tmpl", nil)
}

func (server *PageServer) Contacts(w http.ResponseWriter, r *http.Request) {
	server.Render(w, "contacts.tmpl", nil)
}

// News Server ////////////////////////
//
type NewsServer struct{
	*PageRenderer
}

func (server *NewsServer) List(w http.ResponseWriter, r *http.Request) {
	queryPage  := r.URL.Query().Get("page")
	if queryPage == "" {
		queryPage = "1"
	}
	pageNum, _ := strconv.Atoi(queryPage)
	PageLimit  := 3

	var news []News
	ReadJSON("config/news.json", &news)

	startPoint := ((pageNum * PageLimit) - 3)
	if startPoint <= 0 {
		startPoint = 0
	}

	endPoint   := pageNum * PageLimit
	news_slice := news[startPoint:endPoint]

	p := NewPagination(len(news), PageLimit, pageNum)

	data := struct {
		News    []News
		Pagination Pagination
	}{
		news_slice,
		*p,
	}

	server.Render(w, "news_list.tmpl", data)
}

func (server *NewsServer) Full(w http.ResponseWriter, r *http.Request) {
	newsName := r.URL.Query().Get("id")

	var news []News
	ReadJSON("config/news.json", &news)

	found := false
	var foundNews News
	for _, item := range news {
		if item.Alias == newsName {
			foundNews = item
			found = true
			break
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		foundNews = News{Title: "Not Found 404"}
	}

	data := struct {
		News    News
	}{
		foundNews,
	}
	server.Render(w, "news_full.tmpl", data)
}


// Pagination ////////////////////////
//
type Pagination struct {
	PerPage     int
	TotalAmount int
	CurrentPage int
	TotalPage   int
}

func NewPagination(totalAmount, perPage, currentPage int) *Pagination {
	if currentPage == 0 {
		currentPage = 1
	}

	n := (totalAmount + perPage - 1) / perPage
	if currentPage > n {
		currentPage = n
	}

	return &Pagination{
		PerPage:     perPage,
		TotalAmount: totalAmount,
		CurrentPage: currentPage,
		TotalPage:   n,
	}
}

type Menu struct {
	Name   string `json:"name"`
	Alias  string `json:"alias"`
}

type News struct {
	Title        string `json:"title"`
	Alias        string `json:"alias"`
	Description  string `json:"description"`
	Content      string `json:"content"`
}

func ReadJSON(fileName string, result interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(result)
}