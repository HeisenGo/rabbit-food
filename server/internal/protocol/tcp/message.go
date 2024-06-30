package tcp

import (
	"server/internal/models/address"
	"server/internal/models/auth"
	"server/internal/models/restaurant/menu"
	"server/internal/models/restaurant/restaurant"
	"server/internal/models/user"
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

type AddOperatorToRestaurantRequest struct {
	OperatorPhoneOrEmail string `json:"phone"` //operator phone
	RestaurantID         uint   `json:"restaurant_id"`
}

type CreateRestaurantRequest struct {
	Name    string           `json:"name"`
	Phone   string           `json:"phone"`
	Address *address.Address `json:"address"`
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

type GetRestaurantCategoriesRequest struct {
	RestaurantID uint `json:"restaurant_id"`
}

type AddCategoryToRestaurantRequest struct {
	RestaurantID uint   `json:"restaurant_id"`
	CategoryIDs  []uint `json:"category_ids"`
}

type DepositRequest struct {
	CardNumber string `json:"card_number"`
	Amount     uint   `json:"amount"`
}

type WithdrawRequest struct {
	CardNumber string `json:"card_number"`
	Amount     uint   `json:"amount"`
}

type AddressRequest struct {
	AddressLine string              `json:"address_line"`
	Coordinates address.Coordinates `json:"coordinates"`
	City        string              `json:"city"`
}

type AddCardToWalletResponse struct {
	Message string                 `json:"message"`
	Card    *creditCard.CreditCard `json:"card"`
}

type AssignOperatorToRestaurantResponse struct {
	Message                string
	AssignOperatorResponse *AssignOperatorResponse
}

type AssignOperatorResponse struct {
	OperatorPhoneOrEmail string `json:"operator"`
	RestaurantName       string `json:"restaurant_name"`
}

type CreateRestaurantResponse struct {
	Message    string                 `json:"message"`
	Restaurant *restaurant.Restaurant `json:"restaurant"`
}

type CreateMenuResponse struct {
	Message string     `json:"message"`
	Menu    *menu.Menu `json:"menu"`
}

type AddMenuItemToMenuResponse struct {
	Message  string         `json:"message"`
	MenuItem *menu.MenuItem `json:"menu_item"`
}

type GetMenuItemsOfMenuResponse struct {
	Message   string           `json:"message"`
	MenuItems []*menu.MenuItem `json:"menu_items"`
}

type AddCategoriesToRestaurantResponse struct {
	Message    string                 `json:"message"`
	Restaurant *restaurant.Restaurant `json:"restaurant"`
}

type GetRestaurantCategoriesResponse struct {
	Message    string                           `json:"message"`
	Categories []*restaurant.RestaurantCategory `json:"categories"`
}

type GetAllMenusResponse struct {
	Message string       `json:"message"`
	Menus   []*menu.Menu `json:"menus"`
}

type GetUserWalletCardsResponse struct {
	Message string                   `json:"message"`
	Cards   []*creditCard.CreditCard `json:"cards"`
}

type GetOwnerOperatorRestaurantsResponse struct {
	Message     string
	Restaurants []*restaurant.Restaurant
}

type DepositResponse struct {
	Message string         `json:"message"`
	Wallet  *wallet.Wallet `json:"wallet"`
}

type WithdrawResponse struct {
	Message string         `json:"message"`
	Wallet  *wallet.Wallet `json:"wallet"`
}

type GetWalletResponse struct {
	Message string         `json:"message"`
	Wallet  *wallet.Wallet `json:"wallet"`
}

type AddressResponse struct {
	Message string           `json:"message"`
	Address *address.Address `json:"address"`
}

type EditRestaurantNameRequest struct {
	RestaurantID uint   `json:"restaurant_id"`
	NewName      string `json:"new_name"`
}

type RestaurantToAddCategoryMenuFoodResponse struct {
	Message     string
	Restaurants []*restaurant.Restaurant
}

// User Profile Update Requests and Responses

type UpdateFirstNameRequest struct {
	UserID    uint   `json:"user_id"`
	FirstName string `json:"first_name"`
}

type UpdateFirstNameResponse struct {
	Message string     `json:"message"`
	User    *user.User `json:"user"`
}

type UpdateLastNameRequest struct {
	UserID   uint   `json:"user_id"`
	LastName string `json:"last_name"`
}

type UpdateLastNameResponse struct {
	Message string     `json:"message"`
	User    *user.User `json:"user"`
}

type UpdateEmailRequest struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
}

type UpdateEmailResponse struct {
	Message string     `json:"message"`
	User    *user.User `json:"user"`
}

type UpdatePhoneRequest struct {
	UserID uint   `json:"user_id"`
	Phone  string `json:"phone"`
}

type UpdatePhoneResponse struct {
	Message string     `json:"message"`
	User    *user.User `json:"user"`
}

type UpdatePasswordRequest struct {
	UserID   uint   `json:"user_id"`
	Password string `json:"password"`
}

type UpdatePasswordResponse struct {
	Message string     `json:"message"`
	User    *user.User `json:"user"`
}

type DeleteAddressRequest struct {
	UserID    uint `json:"user_id"`
	AddressID uint `json:"address_id"`
}

type DeleteAddressResponse struct {
	Message string `json:"message"`
}

type AddAddressRequest struct {
	UserID      uint                `json:"user_id"`
	AddressLine string              `json:"address_line"`
	Coordinates address.Coordinates `json:"coordinates"`
	Types       string              `json:"types"`
	City        string              `json:"city"`
}

type AddAddressResponse struct {
	Message string           `json:"message"`
	Address *address.Address `json:"address"`
}
