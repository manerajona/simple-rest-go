package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("."))) // Need an index.html
	http.Handle("/favicon", http.NotFoundHandler())  // This is useful because there's not favicon
	log.Fatal(http.ListenAndServe(":8080", nil))
}
