package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("ping", ping)

	server.Run(":8080")
}

func ping(context *gin.Context) {
	context.JSON(http.StatusOK, map[string]any{"res": 200, "msg": "hello world"})
}
