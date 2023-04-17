package utils

import (
	"fmt"
	"os"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/slyxh2/golang-blog/models"
)

func CreateJWTToken(user *models.User) (string, error) {
	key := os.Getenv("ACCESS_TOKEN_SECRET")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"UserName": user.UserName,
			"Id":       user.Id,
		})
	s, err := t.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return s, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string) (string, error) {
	secret := os.Getenv("ACCESS_TOKEN_SECRET")
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	return claims["Id"].(string), nil
}
