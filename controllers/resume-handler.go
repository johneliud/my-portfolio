package controllers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func ResumeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/resume" {
		ServeErrorPage(w, ErrorPage{
			StatusCode: http.StatusNotFound,
			Message:    "Page Not Found",
		})
		return
	}

	tmplPath := filepath.Join("templates", "resume.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		ServeErrorPage(w, ErrorPage{
			StatusCode: http.StatusInternalServerError,
			Message:    "An Unexpected Error Occurred. Try Again Later",
		})
		log.Printf("Error parsing template %s: %v\n", tmplPath, err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		ServeErrorPage(w, ErrorPage{
			StatusCode: http.StatusInternalServerError,
			Message:    "An Unexpected Error Occurred. Try Again Later",
		})
		log.Printf("Error executing template: %v\n", err)
	}
}
