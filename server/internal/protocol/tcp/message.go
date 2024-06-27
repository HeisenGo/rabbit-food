package tcp

import (
	"server/internal/models/address"
	"server/internal/models/auth"
	"server/internal/models/restaurant/menu"
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

type CreateRestaurantRequest struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Coordinates string `json:"coordinates"`
}

type CreateMenuRequest struct {
	RestaurantID uint   `json:"restaurant_id"`
	Name         string `json:"name"`
}

type GetRestaurantMenusRequest struct {
	RestaurantID uint `json:"restaurant_id"`
}

type AddMenuItemToMenuRequest struct {
	MenuID                        uint   `json:"menu_id"`
	Name                          string `json:"name"`
	Price                         uint   `json:"price"`
	PreparationMinutes            uint   `json:"preparation_minutes"`
	CancellationPenaltyPercentage uint   `json:"cancellation_penalty_percentage"`
}

type GetMenuItemsOfMenuRequest struct {
	MenuID uint `json:"menu_id"`
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
type AddressRequest struct{
	AddressLine string `json:"address_line"`
	Cordinates [2]float64	`json:"cordinates"` //coordinates should be changed to the float 64 type`
	City       string	`json:"city"`
}
type AddCardToWalletResponse struct {
	Message string
	Card    *creditCard.CreditCard
}

type CreateRestaurantResponse struct {
	Message    string
	Restaurant *restaurant.Restaurant
}

type CreateMenuResponse struct {
	Message string
	Menu    *menu.Menu
}

type AddMenuItemToMenuResponse struct {
	Message  string
	MenuItem *menu.MenuItem
}

type GetMenuItemsOfMenuResponse struct {
	Message   string
	MenuItems []*menu.MenuItem
}

type GetAllMenusResponse struct {
	Message string
	Menus   []*menu.Menu
}

type GetUserWalletCardsResponse struct {
	Message string                   `json:"message"`
	Cards   []*creditCard.CreditCard `json:"cards"`
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
