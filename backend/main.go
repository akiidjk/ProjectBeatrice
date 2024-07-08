package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func get_status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "online!!!",
	})
}

func main() {
	r := gin.Default()

	// Define the route
	r.GET("/", get_status)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
