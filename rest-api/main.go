package main

import (
	"net/http"
	"strconv"

	"github.com/i-amare/rest-api/db"
	"github.com/i-amare/rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("ping", ping)

	server.GET("vendors", getAllVendors)
	server.GET("vendors/:id", getVendor)
	server.POST("vendors", createVendor)

	server.Run(":3000")
}

func ping(context *gin.Context) {
	response := gin.H{
		"res":     200,
		"message": "Hello World",
	}
	context.JSON(http.StatusOK, response)
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

	v, _ := models.GetVendor(id)
	res := gin.H{
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

func createVendor(context *gin.Context) {
	var vendor models.Vendor
	err := context.ShouldBindJSON(&vendor)

	if err != nil {
		res := gin.H{
			"message": "Error parsing data",
			"data":    vendor,
		}
		context.JSON(http.StatusBadRequest, res)
		return
	}

	vendor.Save()

	res := gin.H{
		"message": "Vendor created",
		"vendor":  vendor,
	}
	context.JSON(http.StatusCreated, res)
}
