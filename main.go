package main

import (
	"log"
	"net/http"
	"github.com/SanderV1992/golang_simple_blog/site"
	"github.com/SanderV1992/golang_simple_blog/page"
	"github.com/SanderV1992/golang_simple_blog/news"
)

const (
	defaultPort = "8080"
)

func main() {
	router := site.Router{}
	renderer, err := site.NewRenderer(
		"templates/*",
		&router,
	)
	if err != nil {
		log.Fatal(err)
	}

	(&page.Server{renderer}).Register(&router)
	(&news.Server{renderer}).Register(&router)

	http.ListenAndServe(":" + defaultPort + "", &router)
}
