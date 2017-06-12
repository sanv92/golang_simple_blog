package main

import (
	"fmt"
	"os"
	"net/http"
	"text/template"
	"encoding/json"
	"errors"
	"math"
	//"strings"
	"strconv"
)

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

var (
	ErrTemplateDoesNotExist = errors.New("The template does not exist.")
)

var tpl = template.Must(template.New("test").Funcs(template.FuncMap{
	"loop": func(n int) []struct{} {
		return make([]struct{}, n)
	},
	"add": func(x, y int) int {
		return x + y
	},
}).ParseGlob("templates/*"))

func HomePage(w http.ResponseWriter, r *http.Request) {
	var menu []Menu
	ReadJSON("config/menu.json", &menu)

	data := struct {
		Menu    []Menu
	}{
		menu,
	}
	tpl.ExecuteTemplate(w, "home.tmpl", data)
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	var menu []Menu
	ReadJSON("config/menu.json", &menu)

	data := struct {
		Menu    []Menu
	}{
		menu,
	}
	tpl.ExecuteTemplate(w, "about.tmpl", data)
}

func ContactsPage(w http.ResponseWriter, r *http.Request) {
	var menu []Menu
	ReadJSON("config/menu.json", &menu)

	data := struct {
		Menu    []Menu
	}{
		menu,
	}
	tpl.ExecuteTemplate(w, "contacts.tmpl", data)
}

func NewsList(w http.ResponseWriter, r *http.Request) {
	queryPage  := r.URL.Query().Get("page")
	if queryPage == "" {
		queryPage = "1"
	}
	pageNum, _ := strconv.Atoi(queryPage)
	PageLimit  := 3

	var news []News
	ReadJSON("config/news.json", &news)

	var menu []Menu
	ReadJSON("config/menu.json", &menu)

	startPoint := ((pageNum * PageLimit) - 3)
	if startPoint <= 0 {
		startPoint = 0
	}

	endPoint   := pageNum * PageLimit
	news_slice := news[startPoint:endPoint]

	p := NewPagination(len(news), PageLimit, pageNum)

	data := struct {
		News    []News
		Menu    []Menu
		Pagination Pagination
	}{
		news_slice,
		menu,
		*p,
	}

	tpl.ExecuteTemplate(w, "news_list.tmpl", data)
}

func NewsFull(w http.ResponseWriter, r *http.Request) {
	newsName := r.URL.Query().Get("id")

	var news []News
	ReadJSON("config/news.json", &news)

	var menu []Menu
	ReadJSON("config/menu.json", &menu)

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
		Menu    []Menu
		News    News
	}{
		menu,
		foundNews,
	}
	tpl.ExecuteTemplate(w, "news_full.tmpl", data)
}

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

	n := int(math.Ceil(float64(totalAmount) / float64(perPage)))
	if currentPage > n {
		currentPage = n
	}

	totalPage := int(math.Ceil(float64(totalAmount) / float64(perPage)))

	return &Pagination{
		PerPage:     perPage,
		TotalAmount: totalAmount,
		CurrentPage: currentPage,
		TotalPage:   totalPage,
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
