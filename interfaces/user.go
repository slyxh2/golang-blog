package interfaces

import (
	"context"

	"github.com/slyxh2/golang-blog/models"
)

const (
	CollectionUser = "users"
)

type UserRepository interface {
	Create(context.Context, *models.User) error
}
