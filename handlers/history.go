package handlers

import (
	"net/http"

	"github.com/dannybellion/go-ai-app/models"

	"github.com/gin-gonic/gin"
)

func HistoryHandler(c *gin.Context) {
	history, err := models.GetChatHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve history"})
		return
	}
	c.JSON(http.StatusOK, history)
}