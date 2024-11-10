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
		var form models.ContactForm
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Success: false,
				Error:   "Invalid form data",
			})
			return
		}

		if err := utils.ValidateFormData(&form); err != nil {
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Success: false,
				Error:   err.Error(),
			})
			return
		}

		if err := utils.SendEmail(form); err != nil {
			json.NewEncoder(w).Encode(models.ErrorResponse{
				Success: false,
				Error:   "Failed to send email. Please try again later.",
			})
			return
		}

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
