package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(server *gin.Engine) {
	server.GET("/", ping)

	server.GET("vendors", getAllVendors)
	server.GET("vendors/:id", getVendor)
	server.POST("vendors", createVendor)
	server.PUT("vendors/:id", updateVendor)
	server.DELETE("vendors/:id", deleteVendor)

	server.POST("users", createUser)
	server.GET("users", getAllUsers)

	server.POST("menu", createMenuItem)
}

func ping(context *gin.Context) {
	response := gin.H{
		"res":     200,
		"message": "Hello World",
	}
	context.JSON(http.StatusOK, response)
}
