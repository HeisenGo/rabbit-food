package tcp_service

import (
	"bufio"
	"client/errors"
	"client/protocol/tcp"
	"encoding/json"
	"fmt"
	"net"
)

func SetMethodHeader(header map[string]string, method string) {
	header["method"] = method
}

func SetAuthorizationHeader(header map[string]string) {
	token := GetToken()
	header["Authorization"] = token
}

func ReadResponseFromServer(conn net.Conn) ([]byte, error) {
	reader := bufio.NewReader(conn)
	buffer := make([]byte, 4096)
	n, err := reader.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return nil, err
	}
	return buffer[:n], nil
}

func ResponseErrorProduction(data json.RawMessage) error {
	responseErr, err := tcp.DecodeTCPResponseError(data)
	if err != nil {
		return errors.ErrDecodingServerErrorResponse
	}
	return fmt.Errorf(responseErr.Message)
}
