package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i-amare/rest-api/models"
)

func createMenuItem(context *gin.Context) {
	m, err := parseMenuData(context)
	if err != nil {
		return
	}

	err = m.Save()
	if err != nil {
		res := gin.H{
			"message": err.Error(),
			"error":   err,
		}
		context.JSON(http.StatusInternalServerError, res)
		return
	}

	res := gin.H{
		"message":  "Menu item created",
		"menuItem": m,
	}
	context.JSON(http.StatusOK, res)
}

func parseMenuData(context *gin.Context) (models.MenuItem, error) {
	var m models.MenuItem
	err := context.ShouldBindJSON(&m)
	if err != nil {
		res := gin.H{
			"message": "Error parsing data",
			"data":    m,
			"error":   err.Error(),
			"err":     err,
		}
		context.JSON(http.StatusBadRequest, res)
		return models.MenuItem{}, err
	}

	return m, nil
}
