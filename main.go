package main

import (
	"fmt"
	"log"

	"github.com/dannybellion/go-ai-app/config"
	"github.com/dannybellion/go-ai-app/database"
	"github.com/dannybellion/go-ai-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to SQLite
	database.ConnectDB()

	// Initialize router
	r := gin.Default()
	routes.SetupRoutes(r)

	// Start server
	fmt.Println("Server running on port 8080...")
	log.Fatal(r.Run(":8080"))
}