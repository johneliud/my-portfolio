package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/johneliud/my-portfolio/controllers"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 1 || len(os.Args) > 2 {
		fmt.Println("Usage: 'go run .' OR 'go run . [PORT]'")
		return
	}

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		return
	}

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/static/" {
			http.Error(w, "Access Forbidden", http.StatusForbidden)
			return
		}

		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	})

	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/contact", controllers.FormHandler)
	http.HandleFunc("/api/contact", controllers.ErrorHandler(controllers.FormHandler))

	var port string

	switch len(os.Args) {
	case 1:
		port = ":8080"
	case 2:
		port = os.Args[1]
		portNum, err := strconv.Atoi(port[1:])
		if err != nil {
			fmt.Printf("Failed converting %v to an int: %v", port[1:], err)
			return
		}

		if portNum < 1024 || portNum > 65535 {
			fmt.Println("Invalid port number. Must be in the range of 1024-65535")
			return
		}
		port = ":" + strconv.Itoa(portNum)
	}
	log.Printf("Server running on http://localhost%v\n", port)
	http.ListenAndServe(port, nil)
}
