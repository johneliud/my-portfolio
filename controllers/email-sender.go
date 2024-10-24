package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"gopkg.in/gomail.v2"
)

func SendEmail(form ContactForm) error {
	// Create new message
	message := gomail.NewMessage()
	message.SetHeader("From", os.Getenv("SMTP_FROM"))
	message.SetHeader("To", os.Getenv("SMTP_TO"))
	message.SetHeader("Subject", "Contact Form:"+form.Subject)

	// Create email body
	body := "Name: " + form.Name + "\n"
	body += "Email: " + form.Email + "\n"
	body += "Subject: " + form.Subject + "\n\n"
	body += "Message:\n" + form.Message

	message.SetBody("text/plain", body)

	// Create a dialer
	dialer := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		587,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)
	return dialer.DialAndSend(message)
}

func SendResponse(w http.ResponseWriter, response Response, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
