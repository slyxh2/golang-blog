package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slyxh2/golang-blog/interfaces"
	"github.com/slyxh2/golang-blog/models"
	"github.com/slyxh2/golang-blog/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostController struct {
	pr interfaces.PostRepository
}

func CreatePostController(db *mongo.Database) *PostController {
	pr := repository.NewPostRepository(db)
	return &PostController{
		pr: pr,
	}
}

func (pc *PostController) UploadPost(c *gin.Context) {
	var request interfaces.UploadPostRequest
	c.ShouldBind(&request)
	post := &models.Post{
		Id:     primitive.NewObjectID(),
		Header: request.Header,
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusGone, gin.H{
			"message": err.Error(),
		})
		return
	}
	readfile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusGone, gin.H{
			"message": err.Error(),
		})
		return
	}
	result, err := pc.pr.Upload(c, readfile, post)
	if err != nil {
		c.JSON(http.StatusGone, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusGone, gin.H{
		"result": result,
	})
}
