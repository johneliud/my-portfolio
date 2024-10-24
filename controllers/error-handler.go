package controllers

import (
	"html/template"
	"net/http"
)

type ErrorPage struct {
	StatusCode int
	Message    string
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func ErrorHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				serveErrorPage(w, ErrorPage{
					StatusCode: http.StatusInternalServerError,
					Message:    "An Unexpected Error Occurred. Try Again Later",
				})
			}
		}()
		next(w, r)
	}
}

func serveErrorPage(w http.ResponseWriter, errorPage ErrorPage) {
	tmpl, err := template.ParseFiles("/templates/error.html")
	if err != nil {
		http.Error(w, "An Unexpected Error Occurred. Try Again Later", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(errorPage.StatusCode)

	if err := tmpl.Execute(w, errorPage); err != nil {
		http.Error(w, "Error Rendering Template", http.StatusInternalServerError)
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	serveErrorPage(w, ErrorPage{
		StatusCode: http.StatusNotFound,
		Message:    "Page Not Found",
	})
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	serveErrorPage(w, ErrorPage{
		StatusCode: http.StatusMethodNotAllowed,
		Message:    "Method Not Allowed",
	})
}

func ForbiddenHandler(w http.ResponseWriter, r *http.Request) {
	serveErrorPage(w, ErrorPage{
		StatusCode: http.StatusForbidden,
		Message:    "Access Forbidden",
	})
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	serveErrorPage(w, ErrorPage{
		StatusCode: http.StatusInternalServerError,
		Message:    "An Unexpected Error Occurred. Try Again Later",
	})
}

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	serveErrorPage(w, ErrorPage{
		StatusCode: http.StatusBadRequest,
		Message:    "Bad Request",
	})
}
