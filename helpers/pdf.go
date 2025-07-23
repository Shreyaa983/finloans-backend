package helpers

import (
	"fmt"
	"time"

	"finloans-backend/models"

	"github.com/jung-kurt/gofpdf"
)

func GenerateAndEmailPDF(loan models.LoanApplication) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 14)
	pdf.Cell(40, 10, "Loan Application Details")
	pdf.Ln(12)
	pdf.Cell(40, 10, fmt.Sprintf("User ID: %d", loan.UserID))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Amount: %.2f", loan.Amount))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Tenure: %d months", loan.Tenure))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Status: %s", loan.Status))

	filePath := fmt.Sprintf("public/loan_%d_%d.pdf", loan.UserID, time.Now().Unix())
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		fmt.Println("Error saving PDF:", err)
		return
	}

	// Optionally update DB with path
	// db.Conn.Model(&loan).Update("PdfPath", filePath)

	// Email it
	SendEmailWithAttachment("user@example.com", "Your Loan Application", "Here’s your PDF", filePath)
}
