package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/goTutorialLecture/config/db"
	"github.com/goTutorialLecture/model"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get user function created successfully")
	w.Header().Add("content-type", "application/json")
	var result []model.User
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

	for response.Next(context.Background()) {
		var users model.User
		fmt.Println("Result")
		response.Decode(&users)
		result = append(result, users)
	}

	if err := response.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(w).Encode(result)

}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get profile route is created")
}
