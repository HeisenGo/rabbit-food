package users

type ValidationError struct {
	Field   string
	Message string
}

func newValidationError(field string, msg string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: msg,
	}
}

func (e *ValidationError) Error() string {
	return e.Message
}

var (
	ErrInvalidEmail         = newValidationError("email", "invalid email.")
	ErrInvalidPassword      = newValidationError("password", "invalid password ")
	ErrInvalidPhone         = newValidationError("phone", "invalid phone")
	ErrUserPassDoesNotMatch = newValidationError("phone_or_email", "username and password doesn't match.")
)
