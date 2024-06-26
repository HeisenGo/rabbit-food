package tcp

import (
	"server/internal/models/auth"
	"server/internal/models/restaurant/restaurant"
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

type AddOperatorToRestarantRequest struct {
	OperatorPhoneOrEmail string `json:"phone"` //operator phone
	RestaurantID         uint   `json:"restaurant_id"`
}

type CreateRestaurantRequest struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Coordinates string `json:"coordiantes"`
}

type Coordinates struct {
	X float64
	Y float64
}

type DepositRequest struct {
	CardNumber string `json:"card_number"`
	Amount     uint   `json:"amount"`
}
type WithdrawRequest struct {
	CardNumber string `json:"card_number"`
	Amount     uint   `json:"amount"`
}

type AddCardToWalletResponse struct {
	Message string
	Card    *creditCard.CreditCard
}

type AssignOperatorToRestaurantResponse struct {
	Message                string
	AssignOperatorResponse *AssignOperatorResponse
}

type AssignOperatorResponse struct {
	OperatorPhoneOrEmaile string `json:"operator"`
	RestaurantName        string `json:"restaurant_name"`
}

type CreateRestaurantResponse struct {
	Message    string
	Restaurant *restaurant.Restaurant
}

type GetUserWalletCardsResponse struct {
	Message string
	Cards   []*creditCard.CreditCard
}

type GetOwnerOperatorRestaurantsResponse struct {
	Message     string
	Restaurants []*restaurant.Restaurant
}

type DepositResponse struct {
	Message string
	Wallet  *wallet.Wallet
}
type WithdrawResponse struct {
	Message string
	Wallet  *wallet.Wallet
}

type EditRestarantNameRequest struct {
	RestaurantID uint
	NewName      string
}
