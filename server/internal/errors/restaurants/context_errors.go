package restaurants

type RestaurantError struct {
	Field   string
	Message string
}

func newRestaurantError(field string, msg string) *Error {
	return &Error{
		Field:   field,
		Message: msg,
	}
}

func (e *RestaurantError) Error() string {
	return e.Message
}

var (
	ErrFailedRetrieveID = newRestaurantError("restaurant", "failed to retrieve id from context")
	ErrUserNotAllowed   = newRestaurantError("restaurant", "user not allowed")
)
