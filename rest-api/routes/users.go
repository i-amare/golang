package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i-amare/rest-api/models"
)

func createUser(context *gin.Context) {

}

func parseUserData(context *gin.Context) (models.User, error) {
	var u models.User
	err := context.ShouldBindJSON(&u)
	if err != nil {
		res := gin.H{
			"message": "Error parsing data",
			"data":    u,
		}
		context.JSON(http.StatusBadRequest, res)
		return u, err
	}

	return u, nil
}
