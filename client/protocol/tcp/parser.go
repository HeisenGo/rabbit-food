package tcp

import (
	"client/models"
	"encoding/json"
)

func DecodeRegisterRequest(data []byte) (RegisterRequest, error) {
	var req RegisterRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func EncodeRegisterRequest(req RegisterRequest) ([]byte, error) {
	return json.Marshal(req)
}

func EncodeAddCardReqBody(req *AddCardBody) ([]byte, error) {
	return json.Marshal(req)
}

func EncodeDepositReqBody(req *DepositBody) ([]byte, error) {
	return json.Marshal(req)
}

func EncodeWithdrawReqBody(req *WithdrawBody) ([]byte, error) {
	return json.Marshal(req)
}

func DecodeRegisterResponse(data []byte) (RegisterResponse, error) {
	var res RegisterResponse
	err := json.Unmarshal(data, &res)
	return res, err
}

func DecodeAddCardResponse(data []byte) (*AddCardResponse, error) {
	var res *AddCardResponse
	err := json.Unmarshal(data, &res)
	return res, err
}

func DecodeCreditCard(card []byte) (*models.CreditCard, error) {
	var newCard *models.CreditCard
	err := json.Unmarshal(card, &newCard)
	return newCard, err
}

func DecodeToken(data []byte) (*models.Token, error) {
	var t *models.Token
	err := json.Unmarshal(data, &t)
	return t, err
}

func DecodeLoginRequest(data []byte) (LoginRequest, error) {
	var req LoginRequest
	err := json.Unmarshal(data, &req)
	return req, err
}
func EncodeLoginResponse(res LoginResponse) ([]byte, error) {
	return json.Marshal(res)
}
func DecodeLoginResponse(data []byte) (LoginResponse, error) {
	var req LoginResponse
	err := json.Unmarshal(data, &req)
	return req, err
}

func EncodeLoginReqBody(req *LoginBody) ([]byte, error) {
	return json.Marshal(req)
}

func DecodeGetCardsBodyResponse(data []byte) (GetCardsBodyResponse, error) {
	var resBody GetCardsBodyResponse
	err := json.Unmarshal(data, &resBody)
	return resBody, err
}

func DecodeTCPWalletResponse(data []byte) (GetWalletBodyResponse, error) {
	var res GetWalletBodyResponse
	err := json.Unmarshal(data, &res)
	return res, err
}

func EncodeCreateRestaurantRequest(req CreateRestaurantBody) ([]byte, error) {
	return json.Marshal(req)
}

func DecodeCreateRestaurantResponse(data []byte) (CreateRestaurantResponse, error) {
	var res CreateRestaurantResponse
	err := json.Unmarshal(data, &res)
	return res, err
}

func DecodeGetRestaurantsBodyResponse(data []byte) (GetRestaurantsBodyResponse, error) {
	var resBody GetRestaurantsBodyResponse
	err := json.Unmarshal(data, &resBody)
	return resBody, err
}

func EncodeAddCategoryReqBody(req *RestaurantCategoryBody) ([]byte, error) {
	return json.Marshal(req)
}

func EncodeGetCategoriesOfRestaurantReqBody(req *GetRestaurantsCategoriesBody) ([]byte, error) {
	return json.Marshal(req)
}

func DecodeGetCategoriesRestaurantsBodyResponse(data []byte) (GetCategoriesRestaurantBodyResponse, error) {
	var resBody GetCategoriesRestaurantBodyResponse
	err := json.Unmarshal(data, &resBody)
	return resBody, err
}

func EncodeAddMenuReqBody(req *RestaurantMenuBody) ([]byte, error) {
	return json.Marshal(req)
}

func EncodeGetMenusOfRestaurantReqBody(req *GetRestaurantMenusBody) ([]byte, error) {
	return json.Marshal(req)
}

func DecodeGetRestaurantMenusBodyResponse(data []byte) (GetRestaurantMenusBodyResponse, error) {
	var resBody GetRestaurantMenusBodyResponse
	err := json.Unmarshal(data, &resBody)
	return resBody, err
}

func EncodeAddItemToMenuReqBody(req *AddItemToMenuReqBody) ([]byte, error) {
	return json.Marshal(req)
}

func EncodeGetItemsOfMenuReqBody(req *GetMenuItemsBody) ([]byte, error) {
	return json.Marshal(req)
}

func DecodeGetItemsOfMenuBodyResponse(data []byte) (GetItemsOfMenuBodyResponse, error) {
	var resBody GetItemsOfMenuBodyResponse
	err := json.Unmarshal(data, &resBody)
	return resBody, err
}
