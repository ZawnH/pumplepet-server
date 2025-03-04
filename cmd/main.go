package main

import (
	"log"
	"os"

	"pumplepet-server/internal/routes"
	"pumplepet-server/internal/websocket"
	"pumplepet-server/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	database.ConnectDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// Add CORS middleware
	router.Use(cors.Default())

	// Initialize WebSocket manager
	wsManager := websocket.NewManager()
	go wsManager.Run()

	// Routes
	routes.AuthRoutes(router)
	routes.ChatRoutes(router, wsManager)

	router.Run(":" + port)
}
