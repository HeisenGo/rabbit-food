package users

type Error struct {
	Field   string
	Message string
}

func newError(field string, msg string) *Error {
	return &Error{
		Field:   field,
		Message: msg,
	}
}

func (e *Error) Error() string {
	return e.Message
}

var (
	ErrHashPassword  = newError("password", "Password Hashing failed--> internal server Error")
	ErrWrongPassword = newError("password", "Password is wrong")
)
