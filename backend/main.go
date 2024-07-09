package main

import (
	"backend/config"
	"backend/utils"
	"bufio"
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoLogger, WarningLogger, ErrorLogger = utils.Init()
}

// https://github.com/ollama/ollama/blob/main/docs/api.md

func get_status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "online!!!",
	})
}

func check_api() int {
	resp, err := http.Get(config.URL_API + "api/ps")
	if err != nil {
		ErrorLogger.Println(err)
		return 0
	}
	return resp.StatusCode
}

func create_model() {
	InfoLogger.Println("Creating model...")

	var model = utils.ReadModelfile()

	userData := []byte(`{"name": "beatrice","modelfile": "` + model + `"}`)

	request, error := http.NewRequest("POST", config.URL_API+"api/create", bytes.NewBuffer(userData))
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	if error != nil {
		ErrorLogger.Println(error)
		utils.Check(error)
	}

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		ErrorLogger.Println(error)
		utils.Check(error)
	}

	reader := bufio.NewReader(response.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			ErrorLogger.Println(err)
			break
		}

		result := utils.DecodeJson(line)
		if result != nil {
			InfoLogger.Println(result["status"])
		}
	}

	defer response.Body.Close()
}

func main() {
	r := gin.Default()

	InfoLogger.Println("Starting...")

	if check_api() != 200 {
		ErrorLogger.Println("The ollama API is not running or is not reachable .")
		os.Exit(1)
	}

	create_model()

	InfoLogger.Println("Model created...")

	InfoLogger.Println("Started...")

	// Define the route
	r.GET("/", get_status)

	r.Run("127.0.0.1:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
