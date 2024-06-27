package tcp

import (
	"server/internal/models/address"
	"server/internal/models/auth"
	creditCard "server/internal/models/wallet/credit_card"
	"server/internal/models/wallet/wallet"
)

type RegisterRequest struct {
	Phone    string  `json:"phone"`
	Email    *string `json:"email"`
	Password string  `json:"password"`
}

type RegisterResponse struct {
	Message string      `json:"message"`
	Token   *auth.Token `json:"token"`
}

type LoginRequest struct {
	PhoneOrEmail string `json:"phone_or_email"`
	Password     string `json:"password"`
}

type LoginResponse struct {
	Message string      `json:"message"`
	Token   *auth.Token `json:"token"`
}

type AddCardToWalletRequest struct {
	CardNumber string `json:"card_number"`
}

type DepositRequest struct {
	CardNumber string `json:"card_number"`
	Amount     uint   `json:"amount"`
}
type WithdrawRequest struct {
	CardNumber string `json:"card_number"`
	Amount     uint   `json:"amount"`
}
type AddressRequest struct{
	AddressLine string `json:"address_line"`
	Cordinates string	`json:"cordinates"`
 	Types     string		`json:"types"`
	City       string	`json:"city"`
}
type AddCardToWalletResponse struct {
	Message string
	Card    *creditCard.CreditCard
}

type GetUserWalletCardsResponse struct {
	Message string
	Cards   []*creditCard.CreditCard
}
type DepositResponse struct {
	Message string
	Wallet  *wallet.Wallet
}
type WithdrawResponse struct {
	Message string
	Wallet  *wallet.Wallet
}
type AddressResponse struct{
	Message string
	Address *address.Address
}
