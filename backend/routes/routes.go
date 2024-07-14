package routes

import (
	"backend/logger"
	"backend/ollama"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "online!!!",
	})
}

func GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": "0.1.0",
	})
}

type Message struct {
	Content string
}

func CreateMessage(c *gin.Context) {
	var message Message
	c.BindJSON(&message)
	logger.Debug("Message received: " + message.Content)
	ollama.SendMessage(c, message.Content)
}

func ShowModel(c *gin.Context) {
	model := c.Param("model")
	verbose := strings.ToLower(c.DefaultQuery("verbose", "false"))

	if verbose != "true" && verbose != "false" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "verbose must be true or false"})
		return
	}

	response, return_code := ollama.ShowInfoModel(model, verbose)

	if return_code != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"ollama code": return_code})
	} else {
		c.JSON(http.StatusOK, gin.H{"ollama code": return_code, "response": response})
	}

}
