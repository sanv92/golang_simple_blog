package site

import (
	"net/http"
	"fmt"
)

func Abort(status int, w http.ResponseWriter, r *http.Request) {

	switch status {
	case 404:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 page error")
	case 301:
		//
	case 302:
		//
	}

	return
}
