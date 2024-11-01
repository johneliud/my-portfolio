package controllers

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/johneliud/my-portfolio/models"
	"github.com/johneliud/my-portfolio/utils"
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

func CreateBlogPostHandler(store *utils.BlogStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedHandler(w, r)
			log.Println("Only POST request allowed.")
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			BadRequestHandler(w, r)
			log.Printf("Error reading request body: %v\n", err)
			return
		}

		var createRequest models.CreatePostRequest

		if err := json.Unmarshal(body, &createRequest); err != nil {
			BadRequestHandler(w, r)
			log.Printf("Error parsing JSON: %v\n", err)
			return
		}

		// Validate required fields
		if createRequest.Title == "" || createRequest.Content == "" {
			BadRequestHandler(w, r)
			log.Println("Title and Content fields are required.")
			return
		}

		// Create the blog post
		post := &models.BlogPost{
			Title:   createRequest.Title,
			Content: createRequest.Content,
			Summary: createRequest.Summary,
			Tags:    createRequest.Tags,
		}

		// Save the post
		if err := store.SavePost(post); err != nil {
			InternalServerErrorHandler(w, r)
			log.Printf("Error saving post: %v\n", err)
			return
		}

		// Return the created post
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)
	}
}
