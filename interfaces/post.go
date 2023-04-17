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
	Upload(*gin.Context, multipart.File, *models.Post) (*s3manager.UploadOutput, error)
}

type UploadPostRequest struct {
	Header string         `form:"header" binding:"required"`
	File   multipart.File `form:"post" binding:"required"`
}
