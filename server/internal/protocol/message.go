package protocol

type RegisterRequest struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Success bool
	Message string
	UserID  uint
}
