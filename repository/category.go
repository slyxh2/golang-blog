package repository

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/slyxh2/golang-blog/interfaces"
	"github.com/slyxh2/golang-blog/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type categoryRepository struct {
	database   *mongo.Database
	collection string
}

func NewCategoryRepository(db *mongo.Database) *categoryRepository {
	return &categoryRepository{
		database:   db,
		collection: interfaces.CollectionCategory,
	}
}

func (cr *categoryRepository) GetAll(c *gin.Context) ([]interfaces.GetAllCategoryresponse, error) {
	// collection := cr.database.Collection(cr.collection)
	// collection.Find()
	collection := cr.database.Collection(cr.collection)
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var response []interfaces.GetAllCategoryresponse
	for cursor.Next(c) {
		var category interfaces.GetAllCategoryresponse
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}

		response = append(response, category)
	}
	return response, nil
}

func (cr *categoryRepository) Create(c *gin.Context, category *models.Category) error {
	collection := cr.database.Collection(cr.collection)
	var cat models.Category
	err := collection.FindOne(c, bson.M{"name": category.Name}).Decode(&cat)
	if err == nil {
		return errors.New("The Category Exits")
	}
	_, err = collection.InsertOne(c, category)
	return err
}

func (cr *categoryRepository) Get(c *gin.Context, id string) (models.Category, error) {
	collection := cr.database.Collection(cr.collection)
	var category models.Category
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return category, err
	}
	err = collection.FindOne(c, bson.M{"_id": objID}).Decode(&category)
	if err != nil {
		return category, errors.New("Invilid Category Id")
	}
	return category, nil
}

func (cr *categoryRepository) Delete(c *gin.Context, id string) error {
	collection := cr.database.Collection(cr.collection)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result := collection.FindOneAndDelete(c, bson.M{"_id": objID})
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (cr *categoryRepository) Edit(c *gin.Context, id string, name string) error {
	collection := cr.database.Collection(cr.collection)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.UpdateByID(c, objID, bson.M{"$set": bson.M{"name": name}})
	if err != nil {
		return err
	}
	return nil
}
