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

func DecodeCards(data []json.RawMessage) ([]*models.CreditCard, error) {
	var resBody []*models.CreditCard
	for _, raw_json := range data {
		var card *models.CreditCard
		err := json.Unmarshal(raw_json, card)
		if err != nil {
			return nil, err
		}
		resBody = append(resBody, card)
	}
	return resBody, nil
}
