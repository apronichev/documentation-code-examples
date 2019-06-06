package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "<html><body>Hello, World!</body></html>")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}