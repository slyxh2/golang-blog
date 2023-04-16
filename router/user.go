package router

import (
	"github.com/gin-gonic/gin"
	"github.com/slyxh2/golang-blog/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleUserRouter(router *gin.Engine, db *mongo.Database) {
	uc := controllers.CreateUserController(db)
	router.POST("/sign-up", uc.SignUp)
}
