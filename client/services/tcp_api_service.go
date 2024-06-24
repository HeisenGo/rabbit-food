package services

import (
	"client/models"
	"client/protocol/tcp"
	"client/services/tcp_service"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"
)

// TODO: parsers and json encode, decoders must create

type APIService struct {
	host string
	port string
}

var apiServiceInstance *APIService
var once sync.Once

func GetAPIService(host, port string) *APIService {
	once.Do(func() {
		apiServiceInstance = &APIService{
			host: host,
			port: port,
		}
	})
	return apiServiceInstance
}

func (s *APIService) Register(userData *models.User) (*models.Token, error) {
	// API call here
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		return nil, fmt.Errorf("error connecting to server: %v", err)
	}
	defer conn.Close()

	// Send the message to the server
	registerReq := tcp.RegisterRequest{
		Email:    &userData.Email,
		Phone:    userData.Phone,
		Password: userData.Password,
	}
	encodedRegisterRequest, err := tcp.EncodeRegisterRequest(registerReq)
	if err != nil {
		fmt.Println("Encoding Problem") //:To Do
		time.Sleep(time.Second * 2)
		return nil, err
	}

	header := map[string]string{"method": "POST"}
	err = tcp.SendRequest(conn, "auth/register", header, encodedRegisterRequest)
	//fmt.Println("Data has been sent!")
	if err != nil {
		fmt.Println("Error writing to server:", err)
		time.Sleep(time.Second * 2)

		return nil, err
	}

	// Read the response from the server
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	// _, err = bufio.NewReader(conn).ReadString(' ')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		time.Sleep(time.Second * 2)

		return nil, err
	}
	buffer = buffer[:n]
	response, err := tcp.DecodeTCPResponse(buffer)
	if err != nil {
		fmt.Println("Error decoding response", response)
		//time.Sleep(time.Second * 2)

		return nil, err
	}
	if response.StatusCode != uint(201) {
		//var responseData tcp.ResponseError
		fmt.Println(string(response.Data))
		responseErr, err := tcp.DecodeTCPResponseError(response.Data)
		if err != nil {
			fmt.Println("error in decoding server error")
			//time.Sleep(time.Minute * 2)

			return nil, err
		}
		fmt.Println("Error creating", responseErr.Message)
		//time.Sleep(time.Second * 2)

		return nil, fmt.Errorf(responseErr.Message)
	}
	var responseData tcp.RegisterResponse
	fmt.Println(string(response.Data))

	err = json.Unmarshal(response.Data, &responseData)
	if err != nil {
		fmt.Println("Error in decoding a successful response", err)
		// time.Sleep(time.Second * 2)
		// time.Sleep(time.Minute * 1)

		return nil, err
	}
	var token *models.Token
	err = json.Unmarshal(responseData.Token, &token)
	if err != nil {
		fmt.Println("Error in decoding the token part of data")
		return nil, err
	}
	fmt.Println("Register:", token)
	//fmt.Printf("Server response: %s", response)
	return token, nil
}

// TODO: now its mock data
func (s *APIService) GetWallet(req *models.GetWalletReq) (*models.Wallet, error) {
	return &models.Wallet{
		ID:      req.ID,
		Balance: 50000,
	}, nil
}

func (s *APIService) Login(req *tcp.LoginBody) (*models.Token, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		//return tcp.Token{}, fmt.Errorf("error connecting to server: %v", err)
		return nil, fmt.Errorf("error connecting to server: %v", err)
	}
	defer conn.Close()

	encodedLoginRequest, err := tcp.EncodeLoginReqBody(req)
	if err != nil {
		fmt.Println("Encoding Problem") //:To Do
		time.Sleep(time.Second * 2)
		//return tcp.Token{}, err
		return nil, err
	}
	header := map[string]string{"method": "POST"}
	err = tcp.SendRequest(conn, "auth/login", header, encodedLoginRequest)
	//fmt.Println("Data has been sent!")
	if err != nil {
		fmt.Println("Error writing to server:", err)
		time.Sleep(time.Second * 2)

		//return tcp.Token{}, err
		return nil, err
	}

	// Read the response from the server
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	// _, err = bufio.NewReader(conn).ReadString(' ')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		time.Sleep(time.Second * 2)

		//return tcp.Token{}, err
		return nil, err
	}
	buffer = buffer[:n]
	fmt.Println(string(buffer))
	response, err := tcp.DecodeTCPResponse(buffer)
	fmt.Println(response)
	if err != nil {
		fmt.Println("Error decoding response", response)
		//time.Sleep(time.Second * 2)

		//return tcp.Token{}, err
		return nil, err
	}
	if response.StatusCode != uint(200) {
		//var responseData tcp.ResponseError
		fmt.Println(string(response.Data))
		responseErr, err := tcp.DecodeTCPResponseError(response.Data)
		if err != nil {
			fmt.Println("error in decoding server error")
			//time.Sleep(time.Minute * 2)

			//return tcp.Token{}, err
			return nil, err
		}
		fmt.Println("Error creating", responseErr.Message)
		//time.Sleep(time.Second * 2)

		//return tcp.Token{}, fmt.Errorf(responseErr.Message)
		return nil, fmt.Errorf(responseErr.Message)
	}
	var responseData tcp.LoginResponse
	fmt.Println(string(response.Data))

	err = json.Unmarshal(response.Data, &responseData)
	if err != nil {
		fmt.Println("Error in decoding a successful response", err)
		// time.Sleep(time.Second * 2)
		// time.Sleep(time.Minute * 1)

		//return tcp.Token{}, err
		return nil, err
	}
	var token *models.Token
	err = json.Unmarshal(responseData.AuthToken, token)
	if err != nil {
		fmt.Println("Error in decoding the token part of data")
		//return tcp.Token{}, err
		return nil, err
	}
	fmt.Println("Login:", token)
	//fmt.Printf("Server response: %s", response)
	return token, nil
}

func (s *APIService) AddCard(reqBody *tcp.AddCardBody) (*models.CreditCard, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		return nil, fmt.Errorf("error connecting to server: %v", err)
	}
	defer conn.Close()

	token := tcp_service.GetToken()
	header := map[string]string{"method": "POST", "Authorization": token}
	addCardReqBody := tcp.NewAddCardBody(reqBody.CardNumber)
	encodedAddCardReqBody, err := tcp.EncodeAddCardReqBody(addCardReqBody)
	if err != nil {
		fmt.Println("Encoding Problem") //:To Do
		time.Sleep(time.Second * 2)
		//return tcp.Token{}, err
		return nil, err
	}

	err = tcp.SendRequest(conn, "wallets/cards", header, encodedAddCardReqBody)
	//fmt.Println("Data has been sent!")
	if err != nil {
		fmt.Println("Error writing to server:", err)
		time.Sleep(time.Second * 2)

		//return tcp.Token{}, err
		return nil, err
	}

	// Read the response from the server
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	// _, err = bufio.NewReader(conn).ReadString(' ')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		time.Sleep(time.Second * 2)
		return nil, err
	}
	buffer = buffer[:n]
	fmt.Println(string(buffer))
	response, err := tcp.DecodeTCPResponse(buffer)
	fmt.Println(response)
	if err != nil {
		fmt.Println("Error decoding response", response)
		return nil, err
	}
	if response.StatusCode != uint(200) {
		fmt.Println(string(response.Data))
		responseErr, err := tcp.DecodeTCPResponseError(response.Data)
		if err != nil {
			fmt.Println("error in decoding server error")
			return nil, err
		}
		fmt.Println("Error creating", responseErr.Message)
		return nil, fmt.Errorf(responseErr.Message)
	}
	var addCardResBody *tcp.AddCardResponse
	fmt.Println(string(response.Data))

	err = json.Unmarshal(response.Data, addCardResBody)
	if err != nil {
		fmt.Println("Error in decoding a successful response", err)
		return nil, err
	}
	var newCard *models.CreditCard
	err = json.Unmarshal(addCardResBody.Card, newCard)
	if err != nil {
		fmt.Println("Error in decoding the token part of data")
		//return tcp.Token{}, err
		return nil, err
	}
	fmt.Println("Login:", token)
	//fmt.Printf("Server response: %s", response)
	return newCard, nil
}

func (s *APIService) Logout(req *tcp.LogoutUserReq) error {
	//TODO implement me
	return nil
}
