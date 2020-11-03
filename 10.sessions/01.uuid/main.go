package main

import (
	"fmt"
	"net/http"

	// go get github.com/satori/go.uuid
	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		if id, err := uuid.NewV4(); err != nil {
			cookie = &http.Cookie{
				Name:     "session",
				Value:    id.String(),
				HttpOnly: true, // Secure: true,
				Path:     "/",
			}
			http.SetCookie(w, cookie)
		}
	}
	fmt.Println(cookie)
}
