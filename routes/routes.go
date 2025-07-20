package routes

import (
	"finloans-backend/controllers"
	"finloans-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	protected := api.Group("/auth")
	protected.Use(middlewares.RequireAuth())
	{
		protected.GET("/get-profile", controllers.GetProfile)
	}
}
