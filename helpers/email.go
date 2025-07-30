package helpers

import (
	"fmt"
	"os"

	"github.com/go-gomail/gomail"
)

func sendEmailWithAttachment(toEmail, name, pdfPath string) {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER")) // from .env
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Your Loan Application Summary")
	m.SetBody("text/plain", fmt.Sprintf("Hi %s,\n\nPlease find attached your loan application summary.\n\nThank you!", name))
	m.Attach(pdfPath)

	// Dialer
	port := 587 // Gmail SMTP port
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"))

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent to:", toEmail)
}
