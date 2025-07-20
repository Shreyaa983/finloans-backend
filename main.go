package main

import (
	"finloans-backend/config"
	"finloans-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load()
	// Connect DB
	config.ConnectDB()

	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)

	r.Run(":8080") // run server on port 8080
}
