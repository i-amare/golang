package main

import (
	"net/http"

	"github.com/i-amare/rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("ping", ping)

	server.GET("vendors", getVendors)
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

func getVendors(context *gin.Context) {
	v := models.GetAllVendors()
	context.JSON(http.StatusOK, v)
}

func createVendor(context *gin.Context) {
 var vendor models.Vendor
 context.ShouldBindJSON(&vendor)
}