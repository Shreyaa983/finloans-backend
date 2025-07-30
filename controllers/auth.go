package controllers

import (
	"finloans-backend/config"
	"finloans-backend/models"
	"finloans-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register handles user registration
// It expects a JSON body with name, email, and password fields

// @Summary User Registration
// @Description Registers a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.UserResponse true "User details"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func Register(c *gin.Context) {

	var body models.User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	body.Password = string(hash)

	if err := config.DB.Create(&body).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registered Succesfully", "user": body})
}

// Login handles user login
// It expects a JSON body with email and password fields
// If successful, it returns a JWT token

// @Summary User Login
// @Description Logs in a user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func Login(c *gin.Context) {

	var body models.LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token creation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
