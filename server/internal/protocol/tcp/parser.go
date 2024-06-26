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

func DecodeAddOperatorToRestarantRequest(data []byte) (AddOperatorToRestarantRequest, error) {
	var req AddOperatorToRestarantRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func DecodeGetAllOperatorsOFRestaurant(data []byte)(){
	
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

func EncodeRegisterResponse(res RegisterResponse) ([]byte, error) {
	return json.Marshal(res)
}

func DecodeRegisterResponse(data []byte) (RegisterResponse, error) {
	var res RegisterResponse
	err := json.Unmarshal(data, &res)
	return res, err
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

func DecodeCreateRestaurantRequest(data []byte) (CreateRestaurantRequest, error) {
	var req CreateRestaurantRequest
	err := json.Unmarshal(data, &req)
	return req, err
}

func EncodeCreateRestaurantResponse(res CreateRestaurantResponse) ([]byte, error) {
	return json.Marshal(res)
}

func DecodeEditRestarantNameRequest(data []byte)(EditRestarantNameRequest, error){
	var req EditRestarantNameRequest
	err := json.Unmarshal(data, &req)
	return req, err
}
