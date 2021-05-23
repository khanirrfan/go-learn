package main

import (
	"net/http"

	"github.com/goTutorialLecture/controller"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Database

func main() {
	router := mux.NewRouter()
	// testing api - "curl -v localhost:8080"
	// https://www.example.com/users
	router.HandleFunc("/users", controller.GetUser).Methods("GET")
	router.HandleFunc("/profile", controller.GetProfile).Methods("GET")
	// router.HandleFunc("/posts", controller.GetPosts).Methods("GET")

	http.ListenAndServe(":8080", router)
}
