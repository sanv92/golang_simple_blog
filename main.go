package main

import (
	"log"
	"blog2/site"
	"blog2/page"
	"blog2/news"
	"blog2/serve"
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
