package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i-amare/rest-api/models"
	"github.com/i-amare/rest-api/utils"
)

func createUser(context *gin.Context) {
	u, err := parseUserData(context)
	if err != nil {
		return
	}

	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		res := gin.H{
			"message": "Error parsing password",
		}
		context.JSON(http.StatusInternalServerError, res)
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

func getAllUsers(context *gin.Context) {
	usersArr, err := models.GetAllUsers()
	if err != nil {
		res := gin.H{
			"message":  "Error fetching users",
			"error":    err.Error(),
			"detailed": err,
		}
		context.JSON(http.StatusInternalServerError, res)
		return
	}

	fmt.Println("Printing")

	context.JSON(http.StatusOK, usersArr)
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
