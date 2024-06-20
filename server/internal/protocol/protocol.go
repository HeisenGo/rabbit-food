package protocol

import "encoding/json"

type TCPRequest struct {
	Location string            `json:"location"`
	Header   map[string]string `json:"header"`
	Data     json.RawMessage   `json:"data"`
}

func DecodeTCPRequest(data []byte) (*TCPRequest, error) {
	var req TCPRequest
	err := json.Unmarshal(data, &req)
	return &req, err
}
