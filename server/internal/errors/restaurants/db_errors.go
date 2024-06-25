package restaurants

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
	ErrRestaurantNotFound = newError("restaurant", "restaurant Not Found")
	ErrMismatchedOwner    = newError("restaurant", "owner mismatched")
)
