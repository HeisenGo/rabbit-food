package protocol

type RegisterRequest struct {
	Phone    string
	Email    string
	Password string
}

type RegisterResponse struct {
	Success bool
	Message string
	UserID  uint
}
