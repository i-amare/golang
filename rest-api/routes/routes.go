package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/i-amare/rest-api/middleware"
)

func InitRoutes(server *gin.Engine) {
	server.GET("/", ping)

	loadUserRoutes(server)
	loadVendorRoutes(server)
}

func loadUserRoutes(server *gin.Engine) {
	server.POST("signup", createUser)
	server.POST("login", loginUser)
	server.GET("users", getAllUsers)
}

func loadVendorRoutes(server *gin.Engine) {
	server.GET("vendors", getAllVendors)
	server.GET("vendors/:id", getVendor)
	
	privileged := server.Group("/")
	privileged.Use(middleware.Authenticate)
	privileged.POST("vendors", createVendor)
	privileged.PUT("vendors/:id", updateVendor)
	privileged.DELETE("vendors/:id", deleteVendor)
	privileged.POST("menu", createMenuItem)
}
func ping(context *gin.Context) {
	response := gin.H{
		"res":     200,
		"message": "Hello World",
	}
	context.JSON(http.StatusOK, response)
}
