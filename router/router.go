package router

import (
	"log"
	"os"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(client *mongo.Client) {
	router := gin.Default()
	db := client.Database("blog")
	HandleUserRouter(router, db)
	// router.Run()
	err := endless.ListenAndServe(":8080", router)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 8080 stopped")

	os.Exit(0)
}
