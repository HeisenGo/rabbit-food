package users


var (
	ErrUserNotFound  = newError("user", "User Not Found")
	ErrEmailNotFound = newError("email","Email Not Found")
	ErrPhoneNotFound = newError("phone","Phone Not Found")
)