package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/johneliud/my-portfolio/models"
	"github.com/johneliud/my-portfolio/utils"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
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
		var form models.ContactForm

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

		if err := utils.SendEmail(form); err != nil {
			InternalServerErrorHandler(w, r)
			log.Printf("Failed to send email: %v\n", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.Response{
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
