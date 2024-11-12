package models

import "time"

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

type ErrorPage struct {
	StatusCode int
	Message    string
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type BlogPost struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Content     string    `json:"content"`
	Date        time.Time `json:"date"`
	ReadingTime int       `json:"readingTime"`
	Tags        []string  `json:"tags"`
	ExternalURL string    `json:"externalUrl,omitempty"`
}

type CreatePostRequest struct {
	Title   string `json:"title"`
	Summary string `json:"summary,omitempty"`
	// Content     string   `json:"content"`
	Tags        []string `json:"tags,omitempty"`
	ExternalURL string   `json:"externalUrl,omitempty"`
}
