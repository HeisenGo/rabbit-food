package protocol

import (
	"encoding/json"
)

func DecodeRegisterRequest(data []byte) (RegisterRequest, error) {
	var req RegisterRequest
	err := json.Unmarshal(data, &req)
	return req, err
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
