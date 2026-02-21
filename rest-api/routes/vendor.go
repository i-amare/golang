package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/i-amare/rest-api/models"
)

func createVendor(context *gin.Context) {
	var v models.Vendor
	err := context.ShouldBindJSON(&v)

	if err != nil {
		res := gin.H{
			"message": "Error parsing data",
			"data":    v,
		}
		context.JSON(http.StatusBadRequest, res)
		return
	}

	v.Save()

	res := gin.H{
		"message": "Vendor created",
		"vendor":  v,
	}
	context.JSON(http.StatusCreated, res)
}

func getVendor(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		res := gin.H{
			"message": "Error parsing vendor ID",
		}
		context.JSON(http.StatusBadRequest, res)
		return
	}

	v, err := models.GetVendor(id)
	if err != nil {
		res := gin.H{
			"message": "Vendor not found",
		}
		context.JSON(http.StatusBadRequest, res)
		return
	}

	res := gin.H{
		"message": "Vendor found",
		"vendor":  v,
	}
	context.JSON(http.StatusOK, res)
}

func getAllVendors(context *gin.Context) {
	v, err := models.GetAllVendors()
	if err != nil {
		res := gin.H{
			"message": "Error fetching events",
		}
		context.JSON(http.StatusInternalServerError, res)
		return
	}
	context.JSON(http.StatusOK, v)
}

func updateVendor(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		res := gin.H{
			"message": "Error parsing vendor ID",
		}
		context.JSON(http.StatusBadRequest, res)
		return
	}

	var v = models.Vendor{ID: id}
	err = context.ShouldBindJSON(&v)
	if err != nil {
		res := gin.H{
			"message": "Error parsing data",
			"data":    v,
		}
		context.JSON(http.StatusBadRequest, res)
		return
	}

	_, err = models.GetVendor(id)
	if err != nil {
		res := gin.H{
			"message": "Vendor not found",
		}
		context.JSON(http.StatusBadRequest, res)
		return
	}

	_, err = models.UpdateVendor(v)
	if err != nil {
		res := gin.H{
			"message": "Error updating vendor",
		}
		context.JSON(http.StatusBadRequest, res)
		return
	}

	res := gin.H{
		"message": "Successfully updated vendor",
		"vendor":  v,
	}
	context.JSON(http.StatusAccepted, res)
}

func deleteVendor(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := gin.H{
			"message": "Error parsing vendor ID",
			"data": gin.H{
				"param": context.Params,
			},
		}
		context.JSON(http.StatusBadRequest, res)
		return
	}

	_, err = models.GetVendor(id)
	if err != nil {
		res := gin.H{
			"message": "Error finding vendor",
		}
		context.JSON(http.StatusBadRequest, res)
		return
	}

	err = models.DeleteVendor(id)
	if err != nil {
		res := gin.H{
			"message": "Error deleting database entry",
		}
		context.JSON(http.StatusInternalServerError, res)
		return
	}

	res := gin.H{
		"message": "Vendor deleted",
	}
	context.JSON(http.StatusOK, res)
}
