package tcp

import (
	"encoding/json"
	"net"
)

func SendResponse(conn net.Conn, statusCode uint, header map[string]string, data json.RawMessage) {
	tcpResponse := NewTCPResponse(statusCode, header, data)
	encodeResponse, err := encodeTCPResponse(tcpResponse)
	if err != nil {
		tcpResponse = NewTCPResponse(500, nil, nil)
		encodeResponse, _ = encodeTCPResponse(tcpResponse)
	}
	conn.Write(encodeResponse)
}

func Error(conn net.Conn, statusCode uint, header map[string]string, message string) {
	errData := NewResponseError(message)
	encodedErrData, _ := encodeErrData(errData)
	tcpResponse := NewTCPResponse(statusCode, header, encodedErrData)
	encodeResponse, _ := encodeTCPResponse(tcpResponse)
	conn.Write(encodeResponse)
}

type Request struct {
	Location string            `json:"location"`
	Header   map[string]string `json:"header"`
	Data     json.RawMessage   `json:"data"`
}

type Response struct {
	StatusCode uint              `json:"status_code"`
	Header     map[string]string `json:"header"`
	Data       json.RawMessage   `json:"data"`
}

type ResponseError struct {
	Message string
}

func NewResponseError(message string) *ResponseError {
	return &ResponseError{Message: message}
}

func NewTCPResponse(statusCode uint, header map[string]string, data json.RawMessage) *Response {
	return &Response{
		StatusCode: statusCode,
		Header:     header,
		Data:       data,
	}
}

func DecodeTCPRequest(data []byte) (*Request, error) {
	var req Request
	err := json.Unmarshal(data, &req)
	return &req, err
}

func encodeTCPResponse(res *Response) ([]byte, error) {
	return json.Marshal(res)
}

func encodeErrData(res *ResponseError) ([]byte, error) {
	return json.Marshal(res)
}
