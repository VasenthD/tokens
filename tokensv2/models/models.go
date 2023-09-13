package models

type Info struct {
	Id string `json:"id" bson:"id"`
	Email string `json:"email" bson:"email"`
}