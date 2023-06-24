package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName  string             `json:"firstName,omitempty"`
	MiddleName string             `json:"middleName,omitempty"`
	LastName   string             `json:"lastName,omitempty"`
	Email      string             `json:"email,omitempty"`
	Password   string             `json:"password,omitempty"`
	Role       primitive.ObjectID `json:"role,omitempty"`
}
