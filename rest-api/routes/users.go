package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i-amare/rest-api/models"
)

func createUser(context *gin.Context) {
	u, err := parseUserData(context)
	if err != nil {
		return
	}

	err = u.Save()
	if err != nil {
		res := gin.H{
			"message": err.Error(),
			"error":   err,
		}
		context.JSON(http.StatusInternalServerError, res)
		return
	}

	res := gin.H{
		"message": "User created",
		"vendor":  u,
	}
	context.JSON(http.StatusCreated, res)
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
