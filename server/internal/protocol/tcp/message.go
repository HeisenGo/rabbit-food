package tcp

import "server/internal/models/auth"

type RegisterRequest struct {
	Phone    string  `json:"phone"`
	Email    *string `json:"email"`
	Password string  `json:"password"`
}
type RegisterResponse struct {
	Message string
	Token   *auth.Token
}
type LoginRequest struct {
	PhoneOrEmail string `json:"phone_or_email"`
	Password     string `json:"password"`
}
type LoginResponse struct {
	Message   string
	AuthToken *auth.Token
}
