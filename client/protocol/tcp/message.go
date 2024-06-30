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

type DepositBody struct {
	Number string `json:"card_number"`
	Amount uint   `json:"amount"`
}

type RestaurantCategoryBody struct {
	RestaurantID uint   `json:"restaurant_id"`
	Category_ids []uint `json:"category_ids"`
}

type WithdrawBody struct {
	Number string `json:"card_number"`
	Amount uint   `json:"amount"`
}

type AddCardBody struct {
	CardNumber string `json:"card_number"`
}

type CreateRestaurantBody struct {
	Name    string          `json:"name"`
	Phone   string          `json:"phone"`
	Address *models.Address `json:"address"`
}

type GetCardsBodyResponse struct {
	Message string               `json:"message"`
	Cards   []*models.CreditCard `json:"cards"`
}

type GetWalletBodyResponse struct {
	Message string         `json:"message"`
	Wallet  *models.Wallet `json:"wallet"`
}

type CreateRestaurantRequest struct {
	Name    string          `json:"name"`
	Phone   string          `json:"phone"`
	Address *models.Address `json:"address"`
}

type CreateRestaurantResponse struct {
	Message    string             `json:"message"`
	Restaurant *models.Restaurant `json:"restaurant"`
}

type GetRestaurantsBodyResponse struct {
	Message     string               `json:"message"`
	Restaurants []*models.Restaurant `json:"restaurants"`
}

type GetRestaurantsCategoriesBody struct {
	RestaurantID uint `json:"restaurant_id"`
}

type GetCategoriesRestaurantBodyResponse struct {
	Message    string                       `json:"message"`
	Categories []*models.RestaurantCategory `json:"restaurant_categories"`
}

type RestaurantMenuBody struct {
	RestaurantID uint   `json:"restaurant_id"`
	Name         string `json:"name"`
}

type GetRestaurantMenusBody struct {
	RestaurantID uint `json:"restaurant_id"`
}

type GetRestaurantMenusBodyResponse struct {
	Message string                   `json:"message"`
	Menus   []*models.RestaurantMenu `json:"restaurant_menus"`
}

type AddItemToMenuReqBody struct {
	Name                          string `json:"name"`
	Price                         uint   `json:"price"`
	PreparationMinutes            uint   `json:"preparation_minutes"` // in minutes
	CancellationPenaltyPercentage uint   `json:"cancellation_penalty_percentage"`
	MenuID                        uint   `json:"menu_id"`
}

type GetMenuItemsBody struct {
	MenuID uint `json:"menu_id"`
}

type GetItemsOfMenuBodyResponse struct {
	Message string             `json:"message"`
	Items   []*models.MenuItem `json:"menu_items"`
}

type RestaurantMotorReqBody struct {
	Name         string `json:"name"`
	RestaurantID uint   `json:"restaurant_id"`
	Speed        int    `json:"speed"`
}

type RestaurantAddOperatorReqBody struct {
	PhoneOrEmail string `json:"phone_or_email"`
	RestaurantID uint   `json:"restaurant_id"`
}

func NewAddCardBody(cardNumber string) *AddCardBody {
	return &AddCardBody{CardNumber: cardNumber}
}

type LogoutUserReq struct {
	// TODO
}
