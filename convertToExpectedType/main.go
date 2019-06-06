package main

import (
	"fmt"
	"net/http"
)

func getUserAgent(w http.ResponseWriter, r *http.Request) {
	ua := r.UserAgent()
	fmt.Printf("user agent is: %s \n", ua)
	w.Write("user agent is " + ua)
}

func main() {
	http.HandleFunc("/", getUserAgent)
	http.ListenAndServe(":8080", nil)
}