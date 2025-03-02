package routes

import (
	"github.com/dannybellion/go-ai-app/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/chat", handlers.ChatHandler)
	r.GET("/history", handlers.HistoryHandler)
}