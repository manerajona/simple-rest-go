package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {

	log.Println("Starting the application...")

	/* Mongo*/
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

	/* Mux */
	r := mux.NewRouter()
	r.HandleFunc("/api/message", findAllMessages).Methods("GET")       // findAll
	r.HandleFunc("/api/message/{id}", getMessage).Methods("GET")       // get(id)
	r.HandleFunc("/api/message", createMessage).Methods("POST")        // create
	r.HandleFunc("/api/message/{id}", updateMessage).Methods("PUT")    // update(id)
	r.HandleFunc("/api/message/{id}", deleteMessage).Methods("DELETE") // delete(id)

	log.Fatal(http.ListenAndServe(":8080", r))
}
