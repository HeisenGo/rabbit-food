package tcp

import (
	"server/internal/models/auth"
	creditCard "server/internal/models/wallet/credit_card"
)

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

type AddCardToWalletRequest struct {
	CardNumber string `json:"card_number"`
}

type AddCardToWalletResponse struct {
	Message string
	Card    *creditCard.CreditCard
}
