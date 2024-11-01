package controllers

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/johneliud/my-portfolio/models"
	"github.com/johneliud/my-portfolio/utils"
)

func BlogHandler(store *utils.BlogStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/blog" {
			NotFoundHandler(w, r)
			return
		}

		// Load all blog posts
		posts, err := store.LoadPosts()
		if err != nil {
			log.Printf("Error loading posts: %v\n", err)
			InternalServerErrorHandler(w, r)
			return
		}

		// Create template data
		data := struct {
			Posts []models.BlogPost
		}{
			Posts: posts,
		}

		tmplPath := filepath.Join("templates", "blog.html")
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			InternalServerErrorHandler(w, r)
			log.Printf("Error parsing template %s: %v\n", tmplPath, err)
			return
		}

		if err = tmpl.Execute(w, data); err != nil {
			InternalServerErrorHandler(w, r)
			log.Printf("Error executing template: %v\n", err)
		}
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
			Title:       createRequest.Title,
			Summary:     createRequest.Summary,
			Content:     createRequest.Content,
			Tags:        createRequest.Tags,
			ExternalURL: createRequest.ExternalURL,
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

func SingleBlogPostHandler(store *utils.BlogStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract post ID from URL
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 4 || parts[1] != "blog" || parts[2] != "post" {
			NotFoundHandler(w, r)
			return
		}

		postID := parts[3]

		// Load the specific post
		post, err := store.GetPost(postID)
		if err != nil {
			NotFoundHandler(w, r)
			log.Printf("Error loading post %s: %v\n", postID, err)
			return
		}

		// Parse and execute template
		tmplPath := filepath.Join("templates", "blog-post.html")
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			InternalServerErrorHandler(w, r)
			log.Printf("Error parsing template %s: %v\n", tmplPath, err)
			return
		}

		if err := tmpl.Execute(w, post); err != nil {
			InternalServerErrorHandler(w, r)
			log.Printf("Error executing template: %v\n", err)
			return
		}
	}
}
