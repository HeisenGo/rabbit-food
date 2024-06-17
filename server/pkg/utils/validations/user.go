package validations

import (
	"regexp"
	"server/internal/errors/users"
	"server/internal/models/user"
)

func ValidateUserRegistration(user *user.User) error {
	err := ValidateEmail(user.Email)
	if err != nil {
		return err
	}
	return nil
}

func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	isMatched := emailRegex.MatchString(email)
	if !isMatched {
		return users.ErrInvalidEmail
	}
	return nil
}
