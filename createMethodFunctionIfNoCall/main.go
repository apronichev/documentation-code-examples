package main

import "net/http"

type Handler struct {
}

func main() {
	h := Handler{}
	http.HandleFunc("/", h.handle)
	http.ListenAndServe(":8090", nil)
}
