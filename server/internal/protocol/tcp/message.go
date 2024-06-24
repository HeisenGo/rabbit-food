package tcp

import (
	"server/internal/models/auth"
	"server/internal/models/user"
	"time"
)

type RegisterRequest struct {
	Phone     string  `json:"phone"`
	Email     *string `json:"email"`
	Password  string  `json:"password"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
}

type RegisterResponse struct {
	Message string
	Token   *auth.Token
}

type UserRequest struct {
	ID        uint      `json:"id"`
	Phone     string    `json:"phone"`
	Email     *string   `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
	Address   string    `json:"address"`
}

type UserResponse struct {
	Message string
	User    *user.User
}

type LoginRequest struct {
	PhoneOrEmail string `json:"phone_or_email"`
	Password     string `json:"password"`
}

type LoginResponse struct {
	Message   string
	AuthToken *auth.Token
}
