package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	UserName string             `bson:"username" json:"username"`
	Password []byte             `bson:"password" json:"password"`
}
