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
type LoginRequest struct{
	Phone string `json:"phone"`
	Email string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Success bool
	Message string
	Usertoken *services.UserToken
}
type Token struct{
	Logtoken string
	RefreshToken string
	ExpiresAt         int64
}