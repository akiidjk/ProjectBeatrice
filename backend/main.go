package main

import (
	"backend/logger"
	"backend/ollama"
	"backend/routes"
	"os"

	"github.com/gin-gonic/gin"
)

// https://github.com/ollama/ollama/blob/main/docs/api.md

func main() {
	logger.SetLevel(logger.DebugLevel)

	r := gin.Default()

	logger.Info("Starting...")

	if ollama.CheckApi() != 200 {
		logger.Error("The ollama API is not running or is not reachable .")
		os.Exit(1)
	}

	ollama.CreateModel()

	logger.Success("Model created...")

	logger.Success("Started...")

	// Define the route
	r.GET("/", routes.GetStatus)
	r.GET("/version", routes.GetVersion)
	r.POST("/message", routes.CreateMessage)

	r.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
