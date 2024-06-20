package protocol

import (
	"server/services"
)

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
type LoginRequest struct {
	PhoneOrEmail string `json:"phone_or_email"`
	Password     string `json:"password"`
}
type LoginResponse struct {
	Success   bool
	Message   string
	AuthToken *services.AuthToken
}
type Token struct {
	AuthToken    string
	RefreshToken string
	ExpiresAt    int64
}
