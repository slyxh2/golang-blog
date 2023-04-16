package router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(client *mongo.Client) {
	router := gin.Default()
	db := client.Database("blog")
	HandleUserRouter(router, db)
	router.Run()
}
