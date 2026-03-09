package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/i-amare/rest-api/models"
)

func createVendor(context *gin.Context) {
	userID := context.GetInt64("UserID")

	v, err := parseVendorData(context)
	v.OwnerID = userID
	if err != nil {
		return
	}

	err = v.Save()
	if err != nil {
		res := gin.H{
			"message": err.Error(),
			"error":   err,
		}
		context.JSON(http.StatusInternalServerError, res)
		return
	}

	res := gin.H{
		"message": "Vendor created",
		"vendor":  v,
	}
	context.JSON(http.StatusCreated, res)
}

func getVendor(context *gin.Context) {
	id, err := parseVendorID(context)
	if err != nil {
		return
	}

	v, err := fetchVendor(id, context)
	if err != nil {
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
	id, err := parseVendorID(context)
	if err != nil {
		return
	}

	_, err = fetchVendor(id, context)
	if err != nil {
		return
	}

	v, err := parseVendorData(context)
	if err != nil {
		return
	}

	v.ID = id
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
	id, err := parseVendorID(context)
	if err != nil {
		return
	}

	_, err = fetchVendor(id, context)
	if err != nil {
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

func parseVendorID(context *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := gin.H{
			"message": "Error parsing vendor ID",
			"params":  context.Params,
		}
		context.JSON(http.StatusBadRequest, res)
		return -1, err
	}

	return id, nil
}

func fetchVendor(id int64, context *gin.Context) (models.Vendor, error) {
	var v models.Vendor
	v, err := models.GetVendor(id)
	if err != nil {
		res := gin.H{
			"message": "Error finding vendor",
			"id":      id,
		}
		context.JSON(http.StatusBadRequest, res)
		return v, err
	}

	return v, nil
}

func parseVendorData(context *gin.Context) (models.Vendor, error) {
	var v models.Vendor
	err := context.ShouldBindJSON(&v)
	if err != nil {
		res := gin.H{
			"message": "Error parsing data",
			"data":    v,
			"error":   err.Error(),
		}
		context.JSON(http.StatusBadRequest, res)
		return v, err
	}

	return v, nil
}
