package router

import (
	"github.com/gin-gonic/gin"
	"github.com/slyxh2/golang-blog/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleCategoryRouter(router *gin.RouterGroup, db *mongo.Database) {
	cr := controllers.CreateCategoryController(db)
	router.GET("all-category", cr.GetAllCategory)
	router.POST("create-category", cr.CreateCategory)
}
