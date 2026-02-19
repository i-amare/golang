package main

import (
	"math/rand"
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

func getAllVendors(context *gin.Context) {
	v := models.GetAllVendors()
	context.JSON(http.StatusOK, v)
}

func createVendor(context *gin.Context) {
	randID := strconv.FormatInt(int64(rand.Int()), 16)[:6]
	vendor := models.Vendor{ID: randID}
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
