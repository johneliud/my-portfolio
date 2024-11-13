package utils

import (
    "fmt"
    "os"
    "strconv"
    "github.com/johneliud/my-portfolio/models"
    "gopkg.in/gomail.v2"
)

func SendEmail(form models.ContactForm) error {
    // Get SMTP port from environment variable with fallback
    smtpPort := os.Getenv("SMTP_PORT")
    port, err := strconv.Atoi(smtpPort)
    if err != nil {
        port = 587 // fallback to default port
    }

    // Create new message
    message := gomail.NewMessage()
    message.SetHeader("From", os.Getenv("SMTP_FROM"))
    message.SetHeader("To", os.Getenv("SMTP_TO"))
    message.SetHeader("Subject", "Contact Form: "+form.Subject)

    // Create email body with better formatting
    body := fmt.Sprintf(`
New Contact Form Submission

Name: %s
Email: %s
Subject: %s

Message:
%s
`, form.Name, form.Email, form.Subject, form.Message)

    message.SetBody("text/plain", body)

    // Create a dialer with error handling
    dialer := gomail.NewDialer(
        os.Getenv("SMTP_HOST"),
        port,
        os.Getenv("SMTP_USER"),
        os.Getenv("SMTP_PASS"),
    )

    // Add error logging
    if err := dialer.DialAndSend(message); err != nil {
        return fmt.Errorf("failed to send email: %v", err)
    }

    return nil
}