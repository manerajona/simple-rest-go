package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title,omitempty" bson:"title,omitempty"`
	Body   string             `json:"body,omitempty" bson:"body,omitempty"`
	TypeMs *TypeMessage       `json:"typeMessage,omitempty" bson:"typeMessage,omitempty"`
}

type TypeMessage struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
}

var messages []Message

func findAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := client.Database("testdb").Collection("messages")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "Error": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var message Message
		cursor.Decode(&message)
		messages = append(messages, message)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "Error": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(messages)
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var message Message

	// find
	collection := client.Database("testdb").Collection("messages")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Message{ID: id}).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "Error ": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(message)
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message Message
	json.NewDecoder(r.Body).Decode(&message)
	messages = append(messages, message)

	// save
	collection := client.Database("testdb").Collection("messages")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, message)

	json.NewEncoder(w).Encode(result)
}

func updateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range messages {
		if item.ID.Hex() == params["id"] {
			messages = append(messages[:index], messages[index+1:]...)
			var message Message
			json.NewDecoder(r.Body).Decode(&message)
			messages = append(messages, message)
			json.NewEncoder(w).Encode(message)
			return
		}
	}
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range messages {
		if item.ID.Hex() == params["id"] {
			messages = append(messages[:index], messages[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(messages)
}
