package services

import (
	"client/models"
	"client/protocol/tcp"
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

func (s *APIService) Register(userData *models.User) (tcp.Token, error) {
	// API call here
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		return tcp.Token{}, fmt.Errorf("error connecting to server: %v", err)
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
		return tcp.Token{}, err
	}

	//data :=
	// fmt.Sprintf("1{\"phone\": \"%s\", \"email\": \"%s\", \"password\": \"%s\"}", userData.Phone, userData.Email, userData.Password)
	//_, err = conn.Write([]byte(data))
	err = tcp.SendRequest(conn, "auth/register", map[string]string{}, encodedRegisterRequest)
	//fmt.Println("Data has been sent!")
	if err != nil {
		fmt.Println("Error writing to server:", err)
		time.Sleep(time.Second * 2)

		return tcp.Token{}, err
	}

	// Read the response from the server
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	// _, err = bufio.NewReader(conn).ReadString(' ')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		time.Sleep(time.Second * 2)

		return tcp.Token{}, err
	}
	buffer = buffer[:n]
	response, err := tcp.DecodeTCPResponse(buffer)
	if err != nil {
		fmt.Println("Error decoding response", response)
		//time.Sleep(time.Second * 2)

		return tcp.Token{}, err
	}
	if response.StatusCode != uint(201) {
		//var responseData tcp.ResponseError
		fmt.Println(string(response.Data))
		responseErr, err := tcp.DecodeTCPResponseError(response.Data)
		if err != nil {
			fmt.Println("error in decoding server error")
			//time.Sleep(time.Minute * 2)

			return tcp.Token{}, err
		}
		fmt.Println("Error creating", responseErr.Message)
		//time.Sleep(time.Second * 2)

		return tcp.Token{}, fmt.Errorf(responseErr.Message)
	}
	var responseData tcp.RegisterResponse
	fmt.Println(string(response.Data))

	err = json.Unmarshal(response.Data, &responseData)
	if err != nil {
		fmt.Println("Error in decoding a successful response", err)
		// time.Sleep(time.Second * 2)
		// time.Sleep(time.Minute * 1)

		return tcp.Token{}, err
	}
	var token tcp.Token
	err = json.Unmarshal(responseData.Token, &token)
	if err != nil {
		fmt.Println("Error in decoding the token part of data")
		return tcp.Token{}, err
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

func (s *APIService) Login(req *models.LoginUserReq) (*tcp.Token, error) {
	// API call here
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		//return tcp.Token{}, fmt.Errorf("error connecting to server: %v", err)
		return nil, fmt.Errorf("error connecting to server: %v", err)
	}
	defer conn.Close()

	// Send the message to the server
	LoginReq := tcp.LoginRequest{
		PhoneOrEmail: req.PhoneOrEmail,
		Password:     req.Password,
	}
	encodedLoginRequest, err := tcp.EncodeLoginRequest(LoginReq)
	if err != nil {
		fmt.Println("Encoding Problem") //:To Do
		time.Sleep(time.Second * 2)
		//return tcp.Token{}, err
		return nil, err
	}

	//data :=
	// fmt.Sprintf("1{\"phone\": \"%s\", \"email\": \"%s\", \"password\": \"%s\"}", userData.Phone, userData.Email, userData.Password)
	//_, err = conn.Write([]byte(data))
	err = tcp.SendRequest(conn, "auth/login", map[string]string{}, encodedLoginRequest)
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
	var token tcp.Token
	err = json.Unmarshal(responseData.AuthToken, &token)
	if err != nil {
		fmt.Println("Error in decoding the token part of data")
		//return tcp.Token{}, err
		return nil, err
	}
	fmt.Println("Login:", token)
	//fmt.Printf("Server response: %s", response)
	return &token, nil

}

func (s *APIService) Logout(req *models.LogoutUserReq) error {
	//TODO implement me
	return nil
}
