package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
