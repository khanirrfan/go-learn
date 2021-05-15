package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/goTutorialLecture/controller"
)

var client *mongo.Database

func main() {

	router := mux.NewRouter()
	// https://www.example.com/users
	router.HandleFunc("/users", controller.GetUser).Methods("GET")

	http.ListenAndServe(":8080", router)
}
