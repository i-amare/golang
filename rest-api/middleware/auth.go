package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i-amare/rest-api/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		res := gin.H{
			"message": "No auth token provided",
		}
		context.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	userID, err := utils.VerifyAuthToken(token)
	if err != nil {
		res := gin.H{
			"message": "Not authorised",
			"error":   err.Error(),
		}
		context.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	context.Set("UserID", userID)
	context.Next()
}
