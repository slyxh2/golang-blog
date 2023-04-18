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

type CategoryController struct {
	cr interfaces.CategoryRepository
}

func CreateCategoryController(db *mongo.Database) *CategoryController {
	cr := repository.NewCategoryRepository(db)
	return &CategoryController{
		cr: cr,
	}
}

func (cc *CategoryController) GetAllCategory(c *gin.Context) {
	allCategory, err := cc.cr.GetAll(c)
	if err != nil {
		c.JSON(http.StatusGone, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"categories": allCategory,
	})
}

func (cc *CategoryController) GetCategory(c *gin.Context) {
	id := c.Query("id")
	category, err := cc.cr.Get(c, id)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusGone, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var request interfaces.CreateCategoryRequest
	c.ShouldBind(&request)
	category := &models.Category{
		Id:    primitive.NewObjectID(),
		Name:  request.Name,
		Posts: []models.Post{},
	}
	err := cc.cr.Create(c, category)
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

func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	id := c.Query("id")
	err := cc.cr.Delete(c, id)
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
