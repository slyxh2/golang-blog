package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slyxh2/golang-blog/interfaces"
	"github.com/slyxh2/golang-blog/models"
	"github.com/slyxh2/golang-blog/repository"
	"github.com/slyxh2/golang-blog/utils"
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

func (uc *userController) Login(c *gin.Context) {
	var request interfaces.SignupRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, interfaces.ErrorResponse{Message: err.Error()})
		return
	}
	user, err := uc.ur.GetUserByUsername(c, request.UserName)
	if err != nil {
		c.JSON(http.StatusNotFound, interfaces.ErrorResponse{Message: "User not found with the given email"})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, interfaces.ErrorResponse{Message: "Invalid credentials"})
		return
	}
	token, err := utils.CreateJWTToken(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, interfaces.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok":    true,
		"token": token,
	})

}
