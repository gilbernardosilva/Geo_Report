package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(to, newPassword string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	from := smtpUsername
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	msg := []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: Password Reset\r\n"+
		"\r\n"+
		"Your new password: %s\r\n", from, to, newPassword))

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	fmt.Println("Password reset email sent successfully.")
	return nil
}
