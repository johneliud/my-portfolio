package controllers

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
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
	switch r.Method {
	case http.MethodGet:
		// Render contact form page
		tmplPath := filepath.Join("templates", "contact.html")
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			InternalServerErrorHandler(w, r)
			log.Printf("Failed parsing template %s: %v\n", tmplPath, err)
			return
		}

		if err := tmpl.Execute(w, nil); err != nil {
			InternalServerErrorHandler(w, r)
			log.Printf("Failed executing contact template: %v\n", err)
		}
	case http.MethodPost:
		// Process form submission
		var form ContactForm

		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			InternalServerErrorHandler(w, r)
			log.Printf("Failed to decode request body: %v\n", err)
			return
		}

		if err := form.Validate(); err != nil {
			InternalServerErrorHandler(w, r)
			log.Printf("Form validation failed: %v\n", err)
			return
		}

		if err := SendEmail(form); err != nil {
			InternalServerErrorHandler(w, r)
			log.Printf("Failed to send email: %v\n", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			Success: true,
			Message: "Message sent successfully!",
		})
		log.Printf("Successfully sent email from %s\n", form.Email)
	default:
		MethodNotAllowedHandler(w, r)
		log.Printf("Invalid method %s for /api/contact\n", r.Method)
		return
	}
}
