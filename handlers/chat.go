package handlers

import (
	"net/http"

	"github.com/dannybellion/go-ai-app/models"
	"github.com/dannybellion/go-ai-app/utils"
	"github.com/gin-gonic/gin"
)

func ChatHandler(c *gin.Context) {
	var request struct {
		Prompt string `json:"prompt"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	client := utils.NewHTTPClient()
	responseText, err := client.CallOpenAI(request.Prompt)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to call OpenAI"},
		)
		return
	}

	// Save chat history
	models.SaveChat(request.Prompt, responseText)

	c.JSON(http.StatusOK, gin.H{"response": responseText})
}