package interfaces

import (
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/slyxh2/golang-blog/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionPost = "posts"
)

type PostRepository interface {
	Upload(*gin.Context, multipart.File, *models.Post, string) (*s3manager.UploadOutput, error)
	DownLoad(id string) (string, error)
	Delete(*gin.Context, string) error
	Edit(*gin.Context, string, string, multipart.File) error
	GetOne(c *gin.Context, id string) (GetPostResponse, error)
	GetAll(c *gin.Context, page int, size int, categoryId string) ([]GetAllPostResponse, int, error)
}

type UploadPostRequest struct {
	Header     string         `form:"header" binding:"required"`
	CategoryId string         `form:"categoryId" binding:"required"`
	File       multipart.File `form:"post" binding:"required"`
}

type EditPostRequest struct {
	Id     string `form:"id" binding:"required"`
	Header string `form:"header" binding:"required"`
}

type GetAllPostResponse struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Header   string             `bson:"header" json:"header"`
	Date     time.Time          `bson:"date" json:"date"`
	Category primitive.ObjectID `bson:"category" json:"category"`
}

type GetPostResponse struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Header   string             `bson:"header" json:"header"`
	Date     time.Time          `bson:"date" json:"date"`
	Category primitive.ObjectID `bson:"category" json:"category"`
	Content  string             `bson:"content" json:"content"`
}
