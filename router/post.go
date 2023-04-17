package router

import (
	"github.com/gin-gonic/gin"
	"github.com/slyxh2/golang-blog/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandlePostRouter(router *gin.RouterGroup, db *mongo.Database) {
	pc := controllers.CreatePostController(db)
	router.POST("/upload", pc.UploadPost)
	router.POST("/test-upload", pc.TestUpload)
}