package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendResetEmail(toEmail, token string) error {
	from := os.Getenv("COMMUNICATION_EMAIL")
	password := os.Getenv("COMMUNICATION_EMAIL_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	resetLink := fmt.Sprintf("http://localhost:8080/user/reset-password?token=%s", token)

	subject := "Subject: Reset Password\n"
	body := fmt.Sprintf("Click to reset password:\n\n%s", resetLink)

	msg := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, msg)
}
