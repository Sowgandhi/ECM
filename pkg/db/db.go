package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Table struct {
	Collection *mongo.Collection
}

var NewCollection Table

func ConnectDB() *mongo.Collection {
	var endpoint string
	fmt.Println("Enter the Endpoint of the MongoDB Database: ")
	fmt.Scanf("%s", &endpoint)
	clientOptions := options.Client().ApplyURI(endpoint)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	NewCollection.Collection = client.Database("go_rest_api").Collection("events")
	return NewCollection.Collection

}

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func GetError(err error, w http.ResponseWriter) {
	log.Println(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}
	message, _ := json.Marshal(response)
	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
