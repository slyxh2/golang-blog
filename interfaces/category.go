package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/slyxh2/golang-blog/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionCategory = "category"
)

type CategoryRepository interface {
	Create(*gin.Context, *models.Category) error
	GetAll(*gin.Context) ([]GetAllCategoryresponse, error)
}

type CreateCategoryRequest struct {
	Name string `form:"name" binding:"required"`
}

type GetAllCategoryresponse struct {
	Name string             `bson:"name" json:"name"`
	Id   primitive.ObjectID `bson:"_id" json:"id"`
}
