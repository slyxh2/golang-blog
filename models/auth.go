package models

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtCustomClaims struct {
	UserName string             `json:"username"`
	Id       primitive.ObjectID `json:"id"`
	jwt.Claims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.Claims
}
