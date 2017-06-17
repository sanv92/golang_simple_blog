package main

import (
	"log"
	"net/http"
	"golang_simple_blog/site"
	"golang_simple_blog/page"
	//"golang_simple_blog/news"
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

	mux := http.NewServeMux()

	(&page.Server{renderer}).Register(mux, &router)
	//(&news.Server{renderer}).Register(mux, &router)

	http.ListenAndServe(":" + defaultPort + "", mux)
}
