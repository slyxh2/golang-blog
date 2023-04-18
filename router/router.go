package router

import (
	"log"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/slyxh2/golang-blog/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(client *mongo.Client) {
	router := gin.Default()
	db := client.Database("blog")
	HandleUserRouter(router, db)

	protectedRouter := router.Group("")
	protectedRouter.Use(middleware.JwtMiddleware())
	HandlePostRouter(protectedRouter, db)
	HandleCategoryRouter(protectedRouter, db)
	protectedRouter.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := endless.ListenAndServe(":8080", router)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 8080 stopped")

	os.Exit(0)
}
