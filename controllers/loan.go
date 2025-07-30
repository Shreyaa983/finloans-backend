package controllers

import (
	"fmt"
	"net/http"

	"finloans-backend/config"
	"finloans-backend/helpers"
	"finloans-backend/models"

	"github.com/gin-gonic/gin"
)

// CheckEligibility godoc
// @Summary Check loan eligibility
// @Description Check if the user is eligible for a loan based on income and age
// @Tags Loans
// @Accept json
// @Produce json
// @Param eligibility body models.EligibilityRequest true "Eligibility input"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /check-eligibility [post]
// @Security BearerAuth
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

// ApplyLoan godoc
// @Summary Apply for a new loan
// @Description Submits a loan application for the authenticated user
// @Tags Loans
// @Accept json
// @Produce json
// @Param loan body models.LoanRequest true "Loan details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /apply-loan [post]
// @Security BearerAuth
func ApplyLoan(c *gin.Context) {
	var req struct {
		Amount float64 `json:"amount"`
		Tenure int     `json:"tenure"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.MustGet("userID").(uint) // from middleware

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

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

	go helpers.GenerateAndEmailPDF(loan, user.Email, user.Name)

	c.JSON(http.StatusOK, gin.H{"message": "Application submitted", "loan_id": loan.ID})
}

// GetMyLoans godoc
// @Summary Get all loans for the authenticated user
// @Description Fetches all loan applications for the authenticated user
// @Tags Loans
// @Produce json
// @Success 200 {array} models.LoanResponse
// @Failure 500 {object} map[string]interface{}
// @Router /my-loans [get]
// @Security BearerAuth
func GetMyLoans(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	fmt.Println("UserID from getmyloans:", userID)
	var loans []models.LoanApplication

	if err := config.DB.Where("user_id = ?", userID).Find(&loans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch loans"})
		return
	}

	c.JSON(http.StatusOK, loans)
}
