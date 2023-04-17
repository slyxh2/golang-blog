package repository

import (
	"context"

	"github.com/slyxh2/golang-blog/interfaces"
	"github.com/slyxh2/golang-blog/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	database   *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database) *userRepository {
	return &userRepository{
		database:   db,
		collection: interfaces.CollectionUser,
	}
}

func (ur *userRepository) Create(ctx context.Context, user *models.User) error {
	collection := ur.database.Collection(ur.collection)
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (ur *userRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user models.User
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return user, err
}
