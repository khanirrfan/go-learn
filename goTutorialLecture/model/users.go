package model

type User struct {
	Firstname string `json:"firstname, omitempty" bson:"firstname, omitempty"`
	Lastname  string `json:"lastname, omitempty" bson:"lastname, omitempty"`
}
