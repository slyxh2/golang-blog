package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(client *mongo.Client) {
	router := gin.Default()
	db := client.Database("blog")
	fmt.Println("DDDDDBBBB", db.Collection("aa").Name())
	// repository.NewUserRepository(db)
	// router.GET("/ping", controllers.TestController)
	HandleUserRouter(router, db)
	router.Run()
}
