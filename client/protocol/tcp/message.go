package tcp

import (
	"encoding/json"
)

type RegisterRequest struct {
	Phone    string  `json:"phone"`
	Email    *string `json:"email"`
	Password string  `json:"password"`
}

type RegisterResponse struct {
	Message string          `json:"message"`
	Token   json.RawMessage `json:"token"`
}
type LoginRequest struct {
	PhoneOrEmail string `json:"phone_or_email"`
	Password     string `json:"password"`
}
type LoginResponse struct {
	Message   string
	AuthToken json.RawMessage
}

type AddCardResponse struct {
	Message string          `json:"message"`
	Card    json.RawMessage `json:"card"`
}

type LoginUserReq struct {
	Header map[string]string
	Body   *LoginBody
}

type LoginBody struct {
	PhoneOrEmail string
	Password     string
}

type AddCardBody struct {
	CardNumber string `json:"card_number"`
}

func NewAddCardBody(cardNumber string) *AddCardBody {
	return &AddCardBody{CardNumber: cardNumber}
}

type LogoutUserReq struct {
	// TODO
}
