package main

import (
	"log"
	"net/http"
	"github.com/SanderV1992/golang_simple_blog/site"
	"github.com/SanderV1992/golang_simple_blog/page"
	"github.com/SanderV1992/golang_simple_blog/news"

	"github.com/SanderV1992/golang_simple_blog/database"
)

const (
	defaultPort = "8080"
)

func main() {
	DB := database.Connect()

	router := site.Router{}
	renderer, err := site.NewRenderer(
		"templates/*",
		&router,
	)
	if err != nil {
		log.Fatal(err)
	}

	newsRepo := &news.RepoMysql{DB}

	(&page.Server{renderer}).Register(&router)
	(&news.Server{renderer, newsRepo}).Register(&router)

	http.ListenAndServe(":" + defaultPort + "", &router)
}
