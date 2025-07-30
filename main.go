package main

import (
	"finloans-backend/config"
	"finloans-backend/routes"
	_ "finloans-backend/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Finloans Backend API
// @version 1.0
// @description API documentation for Finloans backend
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Load environment variables
	godotenv.Load()
	// Connect DB
	config.ConnectDB()

	r := gin.Default()

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes
	routes.RegisterRoutes(r)

	r.Run(":8080") // run server on port 8080
}
