package tcp

import (
	"client/models"
	"encoding/json"
	"fmt"
	"net"
)

func SendRequest(conn net.Conn, location string, header map[string]string, data json.RawMessage) error {
	request := Request{}
	request.Data = data
	request.Location = location
	request.Header = header
	encodedRequest, err := encodeTCPRequest(&request)
	if err != nil {
		fmt.Println("Problem in encoding request")
		return err
	}
	_, err = conn.Write(encodedRequest)
	return err
}

// func Error(conn net.Conn, statusCode uint, header map[string]string, message string) {
// 	errData := NewResponseError(message)
// 	encodedErrData, _ := encodeErrData(errData)
// 	tcpResponse := NewTCPResponse(statusCode, header, encodedErrData)
// 	encodeResponse, _ := encodeTCPResponse(tcpResponse)
// 	conn.Write(encodeResponse)
// }

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
	Message string `json:"Message"`
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

func DecodeTCPResponse(data []byte) (*Response, error) {
	var resp Response
	err := json.Unmarshal(data, &resp)
	return &resp, err
}

func encodeTCPRequest(res *Request) ([]byte, error) {
	return json.Marshal(res)
}

func DecodeTCPResponseError(data []byte) (*ResponseError, error) {
	var resp ResponseError
	err := json.Unmarshal(data, &resp)
	return &resp, err
}

