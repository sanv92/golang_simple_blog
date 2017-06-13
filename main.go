package main

import (
	"fmt"
	"net/http"
	"log"

	. "blog2/app/site"
	. "blog2/app/page"
	. "blog2/app/news"
)

func main() {
	renderer, err := NewRenderer("templates/*")
	if err != nil {
		log.Fatal(err)
	}
	pages := &PageServer{renderer}
	news  := &NewsServer{renderer}

	http.HandleFunc("/", pages.Home)
	http.HandleFunc("/about", pages.About)
	http.HandleFunc("/contacts", pages.Contacts)

	http.HandleFunc("/news/", news.List)
	http.HandleFunc("/news/show/", news.Full)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	fmt.Println("Listen on port: 8080")
	http.ListenAndServe(":8080", nil)
}

