package controllers

import (
	"net/http"

	"finloans-backend/config"
	"finloans-backend/helpers"
	"finloans-backend/models"

	"github.com/gin-gonic/gin"
)

func CheckEligibility(c *gin.Context) {
	var req struct {
		Income float64 `json:"income"`
		Age    int     `json:"age"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	eligible := req.Income > 25000 && req.Age >= 21 && req.Age <= 60
	reason := "Eligible"
	if !eligible {
		reason = "Does not meet income or age criteria"
	}

	c.JSON(http.StatusOK, gin.H{
		"eligible": eligible,
		"reason":   reason,
	})
}

func ApplyLoan(c *gin.Context) {
	var req struct {
		Amount float64 `json:"amount"`
		Tenure int     `json:"tenure"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.GetInt("userID") // from middleware

	loan := models.LoanApplication{
		UserID:     uint(userID),
		Amount:     req.Amount,
		Tenure:     req.Tenure,
		Status:     "Pending",
		IsEligible: true, // Assume true
		Reason:     "User Applied",
	}

	if err := config.DB.Create(&loan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to apply"})
		return
	}

	go helpers.GenerateAndEmailPDF(loan)

	c.JSON(http.StatusOK, gin.H{"message": "Application submitted", "loan_id": loan.ID})
}


func GetMyLoans(c *gin.Context) {
    userID := c.GetInt("userID")
    var loans []models.LoanApplication

    if err := config.DB.Where("user_id = ?", userID).Find(&loans).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch loans"})
        return
    }

    c.JSON(http.StatusOK, loans)
}
