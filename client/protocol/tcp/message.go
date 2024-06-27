package tcp

import (
	"client/models"
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
	Message   string          `json:"message"`
	AuthToken json.RawMessage `json:"token"`
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
	PhoneOrEmail string `json:"phone_or_email"`
	Password     string `json:"password"`
}

type AddCardBody struct {
	CardNumber string `json:"card_number"`
}

type GetCardsBodyResponse struct {
	Message string            `json:"message"`
	Cards   []*models.CreditCard `json:"cards"`
}

func NewAddCardBody(cardNumber string) *AddCardBody {
	return &AddCardBody{CardNumber: cardNumber}
}

type LogoutUserReq struct {
	// TODO
}
