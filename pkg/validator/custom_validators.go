package validator

import (
	"regexp"
	"unicode"
)

// ValidateUsername validates username format and requirements
func ValidateUsername(username string) error {
	if len(username) < 3 {
		return NewValidationError("username must be at least 3 characters long")
	}

	if len(username) > 50 {
		return NewValidationError("username must be no more than 50 characters long")
	}

	// Username should contain only alphanumeric characters and underscores
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !usernameRegex.MatchString(username) {
		return NewValidationError("username can only contain letters, numbers, and underscores")
	}

	return nil
}

// ValidatePassword validates password strength
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return NewValidationError("password must be at least 8 characters long")
	}

	if len(password) > 100 {
		return NewValidationError("password must be no more than 100 characters long")
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return NewValidationError("password must contain at least one uppercase letter")
	}

	if !hasLower {
		return NewValidationError("password must contain at least one lowercase letter")
	}

	if !hasNumber {
		return NewValidationError("password must contain at least one number")
	}

	if !hasSpecial {
		return NewValidationError("password must contain at least one special character")
	}

	return nil
}

// ValidateEmail validates email format
func ValidateEmail(email string) error {
	if len(email) == 0 {
		return NewValidationError("email is required")
	}

	if len(email) > 255 {
		return NewValidationError("email must be no more than 255 characters long")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return NewValidationError("invalid email format")
	}

	return nil
}

// ValidationError represents a validation error
type ValidationError struct {
	Message string
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	return e.Message
}

// NewValidationError creates a new validation error
func NewValidationError(message string) *ValidationError {
	return &ValidationError{Message: message}
}
