package main

import (
	"fmt"
	"net/http"

	//"text/template"
	//"log"
	//"encoding/json"

	. "blog/pkg/app"
	"blog/pkg/app/json"
)

type Data struct {
	Menu    interface{}
	Content interface{}
}

func middlewareMenu(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var menu []json.Menu
		json.ReadJSON(&menu)

		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", middlewareMenu(HomePage))
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/contacts", ContactsPage)
	http.HandleFunc("/news/", NewsList)
	http.HandleFunc("/news/show/", NewsFull)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	fmt.Println("Listen on port: 8080")
	http.ListenAndServe(":8080", nil)
}

// posterID := strings.SplitN(req.URL.Path, "/", 3)[2]
