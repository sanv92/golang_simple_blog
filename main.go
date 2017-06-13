package main

import (
	"log"
	"golang_simple_blog/site"
	"golang_simple_blog/page"
	"golang_simple_blog/news"
	"golang_simple_blog/serve"
)

const (
	defaultPort = "8080"
)

func main() {
	renderer, err := site.NewRenderer("templates/*")
	if err != nil {
		log.Fatal(err)
	}
	pages     := &page.Server{renderer}
	lastNews  := &news.Server{renderer}

	serve.Handler(pages, lastNews)
	serve.Run(defaultPort)
}
