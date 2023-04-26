package controllers

import (
	"net/http"
	"strconv"
	"time"

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
	pr, err := repository.NewPostRepository(db)
	if err != nil {
		return nil
	}
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
		Date:   time.Now(),
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
	result, err := pc.pr.Upload(c, readfile, post, request.CategoryId)
	if err != nil {
		c.JSON(http.StatusGone, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func (pc *PostController) GetPost(c *gin.Context) {
	postId := c.Query("id")
	content, err := pc.pr.DownLoad(postId)
	if err != nil {
		c.JSON(http.StatusGone, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"content": content,
	})
}

func (pc *PostController) DeletePost(c *gin.Context) {
	postId := c.Query("id")
	err := pc.pr.Delete(c, postId)
	if err != nil {
		c.JSON(http.StatusGone, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (pc *PostController) EditPost(c *gin.Context) {
	var request interfaces.EditPostRequest
	c.ShouldBind(&request)
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
	err = pc.pr.Edit(c, request.Id, request.Header, readfile)
	if err != nil {
		c.JSON(http.StatusGone, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func (pc *PostController) GetOnePost(c *gin.Context) {
	postId := c.Query("id")
	post, err := pc.pr.GetOne(c, postId)
	if err != nil {
		c.JSON(http.StatusGone, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func (pc *PostController) GetAllPost(c *gin.Context) {
	categoryId := c.Query("category")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusGone, gin.H{
			"message": err.Error(),
		})
		return
	}
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		c.JSON(http.StatusGone, gin.H{
			"message": err.Error(),
		})
		return
	}
	posts, totalPage, err := pc.pr.GetAll(c, page, size, categoryId)
	if err != nil {
		c.JSON(http.StatusGone, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts":      posts,
		"totalPage":  totalPage,
		"categoryId": categoryId,
	})
}
