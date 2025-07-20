package controllers

import (
	"net/http"

	"finloans-backend/config"
	"finloans-backend/models"

	"github.com/gin-gonic/gin"
)

// GetProfile retrieves the authenticated user's profile
func GetProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Password = "" // Exclude password from response
	c.JSON(http.StatusOK, user)
}
