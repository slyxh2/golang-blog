package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Id    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Posts []Post             `bson:"posts" json:"posts"`
}
