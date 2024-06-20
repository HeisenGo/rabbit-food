package users

import (
	"regexp"
	"server/internal/errors/users"
	"strings"
)

func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	isMatched := emailRegex.MatchString(email)
	if !isMatched {
		return users.ErrInvalidEmail
	}
	return nil
}

func ValidatePasswordWithFeedback(password string) error {
	tests := []struct {
		pattern string
		message string
	}{
		{".{7,}", "Password must be at least 7 characters long"},
		{"[a-z]", "Password must contain at least one lowercase letter"},
		{"[A-Z]", "Password must contain at least one uppercase letter"},
		{"[0-9]", "Password must contain at least one digit"},
		{"[^\\d\\w]", "Password must contain at least one special character"},
	}

	var errMessages []string

	for _, test := range tests {
		match, _ := regexp.MatchString(test.pattern, password)
		if !match {
			errMessages = append(errMessages, test.message)
		}
	}

	if len(errMessages) > 0 {
		feedback := strings.Join(errMessages, "\n")
		// err := errors.Join(users.ErrInvalidPassword, fmt.Errorf(feedback))
		users.ErrInvalidPassword.Message = users.ErrInvalidPassword.Message + feedback
		return users.ErrInvalidPassword
	}

	return nil
}

func ValidatePhone(phone string) error {
	tests := []string{"^(?:\\+989|09)\\d{9}$"}
	// "^09\\d{9}$",  // just 099...
	possible := false
	for _, test := range tests {
		match, _ := regexp.MatchString(test, phone)
		if match {
			possible = true
		}
	}
	if !possible {
		return users.ErrInvalidPhone
	}
	return nil
}
