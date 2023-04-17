package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id     primitive.ObjectID `bson:"_id" json:"id"`
	Header string             `bson:"header" json:"header"`
}
