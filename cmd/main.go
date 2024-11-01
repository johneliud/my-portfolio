package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/johneliud/my-portfolio/controllers"
	"github.com/johneliud/my-portfolio/utils"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 1 || len(os.Args) > 2 {
		fmt.Println("Usage: 'go run .' OR 'go run . [PORT]'")
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v\n", err)
		return
	}

	// Initialize BlogStore
	blogStore, err := utils.NewBlogStore(filepath.Join(".", "blog-posts"))
	if err != nil {
		log.Printf("Error initializing blog store: %v\n", err)
		return
	}

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/static/" {
			controllers.AccessForbiddenHandler(w, r)
			log.Printf("Access forbidden to /static/\n")
			return
		}
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	})

	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/projects", controllers.ProjectsHandler)
	http.HandleFunc("/skills", controllers.SkillsHandler)
	http.HandleFunc("/blog", controllers.BlogHandler)
	http.HandleFunc("/blog/create", controllers.CreateBlogPostHandler(blogStore))
	http.HandleFunc("/contact", controllers.ContactHandler)
	http.HandleFunc("/resume", controllers.ResumeHandler)

	port := utils.GetPort()
	log.Printf("Server starting on http://localhost%v", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Printf("Server failed to start: %v", err)
		os.Exit(1)
	}
}
