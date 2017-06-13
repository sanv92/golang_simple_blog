package serve

import (
	"net/http"
	"golang_simple_blog/page"
	"golang_simple_blog/news"
)


func Handler(pages *page.Server, lastNews *news.Server) {
	http.HandleFunc("/", pages.Home)
	http.HandleFunc("/about", pages.About)
	http.HandleFunc("/contacts", pages.Contacts)

	http.HandleFunc("/news/", lastNews.List)
	http.HandleFunc("/news/show/", lastNews.Full)

	http.Handle("/favicon.ico", http.NotFoundHandler())
}
