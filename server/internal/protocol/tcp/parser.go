package tcp

import (
	"encoding/json"
)

func DecodeRegisterRequest(data []byte) (RegisterRequest, error) {
	var req RegisterRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func DecodeAddCardToWalletRequest(data []byte) (AddCardToWalletRequest, error) {
	var req AddCardToWalletRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func DecodeAddOperatorToRestaurantRequest(data []byte) (AddOperatorToRestaurantRequest, error) {
	var req AddOperatorToRestaurantRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func DecodeGetAllOperatorsOFRestaurant(data []byte) {

}

func DecodeDepositRequest(data []byte) (DepositRequest, error) {
	var req DepositRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func DecodeWithdrawRequest(data []byte) (WithdrawRequest, error) {
	var req WithdrawRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func EncodeAddCardToWalletResponse(res AddCardToWalletResponse) ([]byte, error) {
	return json.Marshal(res)
}

func EncodeAssignOperatorResponse(res AssignOperatorToRestaurantResponse) ([]byte, error) {
	return json.Marshal(res)
}

func EncodeGetUserWalletCardsResponse(res GetUserWalletCardsResponse) ([]byte, error) {
	return json.Marshal(res)
}

func EncodeGetOwnerOperatorRestaurantsResponse(res GetOwnerOperatorRestaurantsResponse) ([]byte, error) {
	return json.Marshal(res)
}

func EncodeDepositResponse(res DepositResponse) ([]byte, error) {
	return json.Marshal(res)
}

func EncodeWithdrawResponse(res WithdrawResponse) ([]byte, error) {
	return json.Marshal(res)
}

func EncodeGetWalletResponse(res GetWalletResponse) ([]byte, error) {
	return json.Marshal(res)
}

func EncodeRegisterResponse(res RegisterResponse) ([]byte, error) {
	return json.Marshal(res)
}

func DecodeLoginRequest(data []byte) (LoginRequest, error) {
	var req LoginRequest
	err := json.Unmarshal(data, &req)
	return req, err
}
func EncodeLoginResponse(res LoginResponse) ([]byte, error) {
	return json.Marshal(res)
}

func DecodeAddAddressToUserRequest(data []byte) (AddressRequest, error) {
	var req AddressRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func EncodeAddAddressToUserResponse(res AddressResponse) ([]byte, error) {
	return json.Marshal(res)
}

func DecodeCreateRestaurantRequest(data []byte) (CreateRestaurantRequest, error) {
	var req CreateRestaurantRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func DecodeCreateMenuRequest(data []byte) (CreateMenuRequest, error) {
	var req CreateMenuRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func DecodeGetRestaurantMenusRequest(data []byte) (GetRestaurantMenusRequest, error) {
	var req GetRestaurantMenusRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func DecodeAddMenuItemToMenuRequest(data []byte) (AddMenuItemToMenuRequest, error) {
	var req AddMenuItemToMenuRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func DecodeGetMenuItemsOfMenuRequest(data []byte) (GetMenuItemsOfMenuRequest, error) {
	var req GetMenuItemsOfMenuRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func DecodeGetRestaurantCategoriesRequest(data []byte) (GetRestaurantCategoriesRequest, error) {
	var req GetRestaurantCategoriesRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func DecodeAddCategoriesToRestaurantRequest(data []byte) (AddCategoryToRestaurantRequest, error) {
	var req AddCategoryToRestaurantRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func EncodeCreateRestaurantResponse(res CreateRestaurantResponse) ([]byte, error) {
	return json.Marshal(res)
}

func DecodeEditRestaurantNameRequest(data []byte) (EditRestaurantNameRequest, error) {
	var req EditRestaurantNameRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func EncodeCreateMenuResponse(res CreateMenuResponse) ([]byte, error) {
	return json.Marshal(res)
}

func EncodeGetAllMenusResponse(res GetAllMenusResponse) ([]byte, error) {
	return json.Marshal(res)
}

func EncodeAddMenuItemToMenuResponse(res AddMenuItemToMenuResponse) ([]byte, error) {
	return json.Marshal(res)
}

func EncodeGetMenuItemsOfMenuResponse(res GetMenuItemsOfMenuResponse) ([]byte, error) {
	return json.Marshal(res)
}
func DecodeRestaurantToAddCategoryMenuFoodResponse(data []byte)(RestaurantToAddCategoryMenuFoodResponse,error){
	var req RestaurantToAddCategoryMenuFoodResponse
	err := json.Unmarshal(data,&req)
	return req,err
}
func EncodeGetRestaurantToAddCategoryMenuFoodResponse(res RestaurantToAddCategoryMenuFoodResponse)([]byte, error){
	return json.Marshal(res)
}
func EncodeAddCategoriesToRestaurantResponse(res AddCategoriesToRestaurantResponse) ([]byte, error) {
	return json.Marshal(res)
}

func EncodeGetRestaurantCategoriesResponse(res GetRestaurantCategoriesResponse) ([]byte, error) {
	return json.Marshal(res)
}
func DecodeAddMotorToRestaurantRequest(data []byte)(AddMotorToRestaurantRequest,error){
	var req AddMotorToRestaurantRequest
	err := json.Unmarshal(data,&req)
	return req,err
}
func EncodeAddMotorToRestaurantResponse(res AddMotorToRestaurantResponse) ([]byte, error) {
	return json.Marshal(res)
}