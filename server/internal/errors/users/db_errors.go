package users

var (
	ErrUserNotFound = newError("user", "User Not Found")
	ErrUserExists   = newError("user", "User already exists.")
)
