package main

import (
	"log"
	"net/http"

	"github.com/Sowgandhi/ECM/pkg/api"
	"github.com/Sowgandhi/ECM/pkg/db"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		api.ShowEvent(w, r)
	case "POST":
		w.WriteHeader(http.StatusCreated)
		api.CreateEvent(w, r)
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		api.UpdateEvent(w, r)
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		api.DeleteEvent(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {

	http.HandleFunc("/", home)
	db.NewCollection.Collection = db.ConnectDB()
	log.Fatal(http.ListenAndServe(":8081", nil))

}
