package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/goTutorialLecture/config/db"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get user function created successfully")
	w.Header().Add("content-type", "application/json")

	db, err := db.GetDBCollection()
	if err != nil {
		log.Fatal(err)
	}
	collection := db.Collection("users")

	ctx, _ := context.WithTimeout(context.TODO(), 50*time.Second)

	response, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal("err")
	}

	defer response.Close(ctx)

}
