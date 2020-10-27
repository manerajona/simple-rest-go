package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	http.HandleFunc("/dog/", d) // accepts /dog/*
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil) // since is nil use default mux
}
