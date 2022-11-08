package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	extractURL := r.URL.Path[len("/hello/"):]
	fmt.Fprintf(w, "%s", extractURL)
}

func main() {
	http.HandleFunc("/hello/", helloHandler)
	http.ListenAndServe("localhost:3000", nil)
}
