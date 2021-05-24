package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/goTutorialLecture/config/db"
	"github.com/goTutorialLecture/model"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get user function created successfully")
	w.Header().Add("content-type", "application/json") //header
	var result []model.User                            //model
	db, err := db.GetDBCollection()                    //database connection
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// collection
	collection := db.Collection("users")
	// context timeout
	ctx, _ := context.WithTimeout(context.TODO(), 50*time.Second)
	// mongo method
	response, err := collection.Find(context.TODO(), bson.M{})
	// error handling
	if err != nil {
		log.Fatal("err")
	}
	// close the context connection
	defer response.Close(ctx)
	// decoding for getting the data from database
	for response.Next(context.Background()) {
		var users model.User
		fmt.Println("Result")
		response.Decode(&users)
		result = append(result, users)
	}
	// error handling for all sort of api type
	if err := response.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	// encoding for every api return response
	json.NewEncoder(w).Encode(result)

}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get profile route is created")
}
