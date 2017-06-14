package main

import (
	"log"
	"net/http"
	"golang_simple_blog/site"
	"golang_simple_blog/page"
	"golang_simple_blog/news"
)

const (
	defaultPort = "8080"
)

func main() {
	renderer, err := site.NewRenderer("templates/*")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	page  := &page.Server{renderer}
	page.Register(mux)

	news  := &news.Server{renderer}
	news.Register(mux)


	http.ListenAndServe(":8080", mux)
}
