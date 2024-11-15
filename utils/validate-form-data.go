package utils

import (
	"errors"
	"regexp"

	"github.com/johneliud/my-portfolio/models"
)

// Validation patterns
var (
	nameRegex  = regexp.MustCompile(`^[a-zA-Z\s]{3,50}$`)
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func ValidateFormData(cf *models.ContactForm) error {
	if !nameRegex.MatchString(cf.Name) {
		return errors.New("invalid name format: must be 3-50 characters long and contain only letters and spaces")
	}

	if !emailRegex.MatchString(cf.Email) {
		return errors.New("invalid email format")
	}

	if len(cf.Subject) > 100 {
		return errors.New("subject must be less than 100 characters")
	}

	if len(cf.Message) < 10 || len(cf.Message) > 1000 {
		return errors.New("message must be between 10 and 1000 characters")
	}
	return nil
}
