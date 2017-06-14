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

	(&page.Server{renderer}).Register(mux)
	(&news.Server{renderer}).Register(mux)

	http.ListenAndServe(":" + defaultPort + "", mux)
}
