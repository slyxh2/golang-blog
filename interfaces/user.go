package interfaces

import (
	"context"

	"github.com/slyxh2/golang-blog/models"
)

const (
	CollectionUser = "users"
)

type SignupRequest struct {
	UserName string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserRepository interface {
	Create(context.Context, *models.User) error
	GetUserByUsername(context.Context, string) (models.User, error)
}
