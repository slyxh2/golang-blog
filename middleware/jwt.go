package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/slyxh2/golang-blog/interfaces"
	"github.com/slyxh2/golang-blog/utils"
)

func JwtMiddleware() gin.HandlerFunc {
	secret := os.Getenv("ACCESS_TOKEN_SECRET")
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Authorization")
		if len(authToken) == 0 {
			c.JSON(http.StatusUnauthorized, interfaces.ErrorResponse{Message: "Lack Authorization Token"})
			c.Abort()
			return
		}
		authorized, err := utils.IsAuthorized(authToken, secret)
		if authorized {
			userID, err := utils.ExtractIDFromToken(authToken)
			if err != nil {
				c.JSON(http.StatusUnauthorized, interfaces.ErrorResponse{Message: err.Error()})
				c.Abort()
				return
			}
			c.Set("x-user-id", userID)
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, interfaces.ErrorResponse{Message: err.Error()})
		c.Abort()
		return
	}
}
