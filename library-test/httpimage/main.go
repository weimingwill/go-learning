package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

	fileServer := http.FileServer(http.Dir("images"))

	http.Handle("/statics/", http.StripPrefix(strings.TrimRight("/statics/", "/"), fileServer))

	http.ListenAndServe(":8080", nil)
}
