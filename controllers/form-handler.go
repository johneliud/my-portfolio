package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
)

type ContactForm struct {
	Name    string
	Email   string
	Subject string
	Message string
}

type Response struct {
	Success bool
	Message string
	Error   string
}

// Validation patterns
var (
	nameRegex  = regexp.MustCompile(`^[a-zA-Z\s]{3,50}$`)
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func (cf *ContactForm) Validate() error {
	if !nameRegex.MatchString(cf.Name) {
		return errors.New("invalid name format: must be 3-50 characters long and contain only letters and spaces")
	}

	if !emailRegex.MatchString(cf.Email) {
		return errors.New("invalid email format")
	}

	if len(cf.Subject) > 100 {
		return errors.New("subject must be less than 100 characters")
	}

	if len(cf.Message) < 10 || len(cf.Message) > 1000 {
		return errors.New("message must be between 10 and 1000 characters")
	}
	return nil
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		MethodNotAllowedHandler(w, r)
		log.Printf("Invalid method %s for /api/contact\n", r.Method)
		return
	}

	var form ContactForm

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		serveErrorPage(w, ErrorPage{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Request",
		})
		log.Printf("Failed to decode request body: %v\n", err)
		return
	}

	if err := form.Validate(); err != nil {
		serveErrorPage(w, ErrorPage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request",
		})
		log.Printf("Form validation failed: %v\n", err)
		return
	}

	if err := SendEmail(form); err != nil {
		serveErrorPage(w, ErrorPage{
			StatusCode: http.StatusInternalServerError,
			Message:    "An Unexpected Error Occurred. Try Again Later",
		})
		log.Printf("Failed to send email: %v\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Message sent successfully!",
	})
	log.Printf("Successfully sent email from %s\n", form.Email)
}
