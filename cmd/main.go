package main

import (
	"log"
	"os"

	"pumplepet-server/pkg/database"
	"pumplepet-server/internal/routes"
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

	// Routes
	routes.AuthRoutes(router)

	router.Run(":" + port)
}

