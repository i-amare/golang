package main

import (
	"github.com/i-amare/rest-api/db"
	"github.com/i-amare/rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.InitRoutes(server)
	server.Run(":3000")
}
