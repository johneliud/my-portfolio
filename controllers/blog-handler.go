package controllers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/blog" {
		NotFoundHandler(w, r)
		return
	}

	tmplPath := filepath.Join("templates", "blog.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		InternalServerErrorHandler(w, r)
		log.Printf("Error parsing template %s: %v\n", tmplPath, err)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		InternalServerErrorHandler(w, r)
		log.Printf("Error executing template: %v\n", err)
	}
}