package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i-amare/rest-api/models"
)

func createMenuItem() {

}

func parseMenuData(context *gin.Context) (models.MenuItem, error) {
	var m models.MenuItem
	err := context.ShouldBindJSON(m)
	if err != nil {
		res := gin.H{
			"message": "Error parsing data",
			"data":    m,
		}
		context.JSON(http.StatusBadRequest, res)
		return models.MenuItem{}, err
	}

	return m, nil
}
