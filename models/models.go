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
	Content     string    `json:"content"`
	Summary     string    `json:"summary"`
	Date        time.Time `json:"date"`
	ReadingTime int       `json:"readingTime"`
	Tags        []string  `json:"tags"`
}

type CreatePostRequest struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Summary string   `json:"summary,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}
