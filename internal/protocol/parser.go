package protocol

import (
	"encoding/json"
)

//func EncodeRegisterRequest(req RegisterRequest) ([]byte, error) {
//	return json.Marshal(req)
//}

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
