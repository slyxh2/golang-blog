package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slyxh2/golang-blog/interfaces"
	"github.com/slyxh2/golang-blog/models"
	"github.com/slyxh2/golang-blog/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	ur interfaces.UserRepository
}

func CreateUserController(db *mongo.Database) *userController {
	ur := repository.NewUserRepository(db)
	return &userController{
		ur: ur,
	}
}

func (uc *userController) CreateUser(c *gin.Context) {
	user := models.User{
		Id: primitive.NewObjectID(),
		// Password: "ooooo",
	}
	err := uc.ur.Create(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusGone, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func (uc *userController) SignUp(c *gin.Context) {
	var request interfaces.SignupRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, interfaces.ErrorResponse{Message: err.Error()})
		return
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfaces.ErrorResponse{Message: err.Error()})
		return
	}
	user := models.User{
		Id:       primitive.NewObjectID(),
		UserName: request.UserName,
		Password: encryptedPassword,
	}
	err = uc.ur.Create(context.Background(), &user)
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
func TestController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"add": "OK",
	})
}
