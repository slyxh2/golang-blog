package interfaces

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/slyxh2/golang-blog/models"
)

const (
	CollectionPost = "posts"
)

type PostRepository interface {
	Upload(*gin.Context, multipart.File, *models.Post, string) (*s3manager.UploadOutput, error)
	DownLoad(id string) (string, error)
	Delete(*gin.Context, string) error
}

type UploadPostRequest struct {
	Header     string         `form:"header" binding:"required"`
	CategoryId string         `form:"categoryId" binding:"required"`
	File       multipart.File `form:"post" binding:"required"`
}
