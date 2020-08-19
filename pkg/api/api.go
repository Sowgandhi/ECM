package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Sowgandhi/ECM/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	//ID       primitive.ObjectID `json: "id,omitempty" bson:"_id,omitempty"`
	Title    string  `json: "title" bson:"title,omitempty"`
	Language string  `json: "language" bson:"language,omitempty"`
	Genre    string  `json:"genre" bson:"genre,omitempty"`
	Time     string  `json: "time" bson:"time,omitempty"`
	Price    string  `json: "price" bson:"price,omitempty"`
	Artist   *Artist `json:"artist" bson:"artist,omitempty" `
}
type Artist struct {
	Name  string `json:"name" bson:"name,omitempty"`
	Image string `json:"image" bson:"image,omitempty"`
}

func ShowEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event Event
	id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		db.GetError(err, w)
		return
	}
	filter := bson.M{"_id": id}
	err = db.NewCollection.Collection.FindOne(context.TODO(), filter).Decode(&event)
	if err != nil {
		db.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(event)

}

//Create
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		db.GetError(err, w)
		return
	}
	result, err := db.NewCollection.Collection.InsertOne(context.TODO(), event)
	if err != nil {
		db.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

//Update
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")
	id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		db.GetError(err, w)
		return
	}
	var event Event
	filter := bson.M{"_id": id}
	err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		db.GetError(err, w)
		return
	}
	//remove bson
	update := bson.D{
		{"$set", bson.D{

			{"title", event.Title},
			{"language", event.Language},
			{"genre", event.Genre},
			{"price", event.Price},
			{"time", event.Time},
			{"artist", bson.D{
				{"name", event.Artist.Name},
				{"image", event.Artist.Image},
			}},
		}},
	}

	err = db.NewCollection.Collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&event)
	if err != nil {
		db.GetError(err, w)
		return
	}
	//event.ID = id
	json.NewEncoder(w).Encode(event)

}

//Delete
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")
	id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		db.GetError(err, w)
		return
	}
	filter := bson.M{"_id": id}
	deleteResult, err := db.NewCollection.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		db.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(deleteResult)
}
