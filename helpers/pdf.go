package helpers

import (
	"fmt"
	"os"
	"time"

	"finloans-backend/models"

	"github.com/jung-kurt/gofpdf"
)

func GenerateAndEmailPDF(loan models.LoanApplication, email string, name string) {
	// Ensure folder exists
	folderPath := "public"
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		fmt.Println("Error creating public folder:", err)
		return
	}

	// Create PDF
	filePath := fmt.Sprintf("%s/loan_%d_%d.pdf", folderPath, loan.UserID, time.Now().Unix())
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Loan Application Summary")
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Hello %s,", name))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Loan ID: %d", loan.ID))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Amount: Rs. %.2f", loan.Amount))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Tenure: %d months", loan.Tenure))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Status: %s", loan.Status))

	if err := pdf.OutputFileAndClose(filePath); err != nil {
		fmt.Println("Error saving PDF:", err)
		return
	}

	fmt.Println("PDF saved:", filePath)

	// Send Email with gomail
	sendEmailWithAttachment(email, name, filePath)
}
