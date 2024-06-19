package services

import (
	"bufio"
	"client/models"
	"errors"
	"fmt"
	"net"
	"sync"
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

func (s *APIService) Register(userData *models.User) error {
	// API call here
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		return errors.New(fmt.Sprintf("Error connecting to server: %v", err))
	}
	defer conn.Close()

	// Send the message to the server
	data := fmt.Sprintf("1{\"phone\": \"%s\", \"email\": \"%s\", \"password\": \"%s\"}", userData.Phone, userData.Email, userData.Password)
	_, err = conn.Write([]byte(data))
	fmt.Println("Data has been sent!")
	if err != nil {
		fmt.Println("Error writing to server:", err)
		return err
	}

	// Read the response from the server
	response, err := bufio.NewReader(conn).ReadString(' ')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return err
	}
	fmt.Printf("Server response: %s", response)
	return nil
}

// TODO: now its mock data
func (s *APIService) GetWallet(req *models.GetWalletReq) (*models.Wallet, error) {
	return &models.Wallet{
		ID:      req.ID,
		Balance: 50000,
	}, nil
}
