package routes

import (
	"backend/ollama"
	"net/http"

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

func CreateMessage(c *gin.Context) {
	return_code := ollama.SendMessage("Oggi cosa hai fatto?")

	if return_code != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"ollama code": return_code})
	} else {
		c.JSON(http.StatusOK, gin.H{"ollama code": return_code})
	}

}
