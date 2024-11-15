package controllers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/johneliud/my-portfolio/models"
)

func serveErrorPage(w http.ResponseWriter, errorPage models.ErrorPage) {
	tmplPath := filepath.Join("templates", "error.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "An Unexpected Error Occurred. Try Again Later", http.StatusInternalServerError)
		log.Printf("Error parsing template %s: %v\n", tmplPath, err)
		return
	}
	w.WriteHeader(errorPage.StatusCode)

	if err := tmpl.Execute(w, errorPage); err != nil {
		http.Error(w, "An Unexpected Error Occurred. Try Again Later", http.StatusInternalServerError)
		log.Printf("Error executing template: %v\n", err)
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	serveErrorPage(w, models.ErrorPage{
		StatusCode: http.StatusNotFound,
		Message:    "Not Found",
	})
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	serveErrorPage(w, models.ErrorPage{
		StatusCode: http.StatusMethodNotAllowed,
		Message:    "Method Not Allowed",
	})
}

func AccessForbiddenHandler(w http.ResponseWriter, r *http.Request) {
	serveErrorPage(w, models.ErrorPage{
		StatusCode: http.StatusForbidden,
		Message:    "Access Forbidden",
	})
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	serveErrorPage(w, models.ErrorPage{
		StatusCode: http.StatusInternalServerError,
		Message:    "An Unexpected Error Occurred. Try Again Later",
	})
}

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	serveErrorPage(w, models.ErrorPage{
		StatusCode: http.StatusBadRequest,
		Message:    "Bad Request",
	})
}
