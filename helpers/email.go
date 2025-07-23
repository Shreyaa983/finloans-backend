package helpers
import (
	"fmt"
)
func SendEmailWithAttachment(to, subject, body, attachmentPath string) {
    // You can use gomail or net/smtp with MIME multipart
    // Or you can integrate SendGrid/Mailgun if needed
    // For demo, keep it simple here or use fmt.Println
    fmt.Println("Email sent to:", to, "with PDF:", attachmentPath)
}
