package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRoleModel struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Label string             `json:"lable,omitempty"`
	Code  string             `json:"code,omitempty"`
}
