package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Header   string             `bson:"header" json:"header"`
	Date     time.Time          `bson:"date" json:"date"`
	Category Category           `bson:"category" json:"category"`
}
