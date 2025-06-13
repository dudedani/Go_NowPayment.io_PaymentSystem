package customer

import (
	"regexp"
	"strings"
)

// Email represents a validated email address value object
type Email struct {
	Address string
}

// emailRegex is a simple email validation regex
// In production, you might want to use a more sophisticated validation library
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// NewEmail creates a new Email value object with validation
func NewEmail(address string) (Email, error) {
	address = strings.TrimSpace(address)
	
	if address == "" {
		return Email{}, ErrEmptyEmail
	}
	
	if !isValidEmail(address) {
		return Email{}, ErrInvalidEmail
	}
	
	// Normalize email to lowercase
	address = strings.ToLower(address)
	
	return Email{Address: address}, nil
}

// isValidEmail validates email format using regex
func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// String returns the email address as a string
func (e Email) String() string {
	return e.Address
}

// GetDomain returns the domain part of the email address
func (e Email) GetDomain() string {
	parts := strings.Split(e.Address, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

// GetUsername returns the username part of the email address
func (e Email) GetUsername() string {
	parts := strings.Split(e.Address, "@")
	if len(parts) == 2 {
		return parts[0]
	}
	return ""
}

// Equals compares two email addresses for equality
func (e Email) Equals(other Email) bool {
	return e.Address == other.Address
}
