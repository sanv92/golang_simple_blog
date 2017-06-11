package app

import (
	//"fmt"

	"net/http"
	. "blog/pkg/app/view"
	//"blog/pkg/app/json"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	//var menu []json.Menu
	//json.ReadJSON(&menu)

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
