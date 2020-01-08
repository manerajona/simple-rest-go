package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) { // implements handlerFunc
	io.WriteString(res, "doggy doggy doggy")
}
func c(res http.ResponseWriter, req *http.Request) { // implements handlerFunc
	io.WriteString(res, "kitty kitty kitty")
}

func main() {
	http.HandleFunc("/dog/", d)                // path dog/something/else/no/matter/what
	http.Handle("/doggy", http.HandlerFunc(d)) // Receives a Handler
	http.HandleFunc("/cat", c)                 //path /cat (Receives a HandlerFunc)
	http.ListenAndServe(":8080", nil)          // You pass nil so it uses the default server mux
}
