package services

import (
	"client/errors"
	"client/models"
	"client/protocol/tcp"
	"client/services/tcp_service"
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

func (s *APIService) MakeNewTCPConnection() (net.Conn, error) {
	// Validate host and port
	if s.host == "" || s.port == "" {
		return nil, fmt.Errorf("host or port is empty")
	}
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", s.host, s.port), 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("error connecting to server: %w", err)
	}
	return conn, nil
}

func (s *APIService) Register(userData *models.User) (*models.Token, error) {
	location := "auth/register"
	header := make(map[string]string)
	methodHeader := tcp.MethodPost
	tcp_service.SetMethodHeader(header, methodHeader)

	conn, err := s.MakeNewTCPConnection()
	if err != nil {
		return nil, errors.ErrConnectionFailed
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
		return nil, errors.ErrEncodingRequest
	}
	err = tcp.SendRequest(conn, location, header, encodedRegisterRequest)
	if err != nil {
		return nil, errors.ErrWritingToServer
	}
	// Read the response from the server
	buffer, err := tcp_service.ReadResponseFromServer(conn)
	if err != nil {
		return nil, errors.ErrReadingResponse
	}

	response, err := tcp.DecodeTCPResponse(buffer)
	if err != nil {
		return nil, errors.ErrDecodingResponse
	}
	if response.StatusCode != tcp.StatusCreated {
		return nil, tcp_service.ResponseErrorProduction(response.Data)
	}

	var responseData tcp.RegisterResponse
	responseData, err = tcp.DecodeRegisterResponse(response.Data)
	if err != nil {
		return nil, errors.ErrDecodingSuccessfulResponse
	}

	var token *models.Token
	token, err = tcp.DecodeToken(responseData.Token)

	if err != nil {
		return nil, errors.ErrDecodingToken
	}
	return token, nil
}

func (s *APIService) Login(req *tcp.LoginBody) (*models.Token, error) {
	location := "auth/login"
	header := make(map[string]string)
	methodHeader := tcp.MethodPost
	tcp_service.SetMethodHeader(header, methodHeader)

	conn, err := s.MakeNewTCPConnection()
	if err != nil {
		return nil, errors.ErrConnectionFailed
	}
	defer conn.Close()

	encodedLoginRequest, err := tcp.EncodeLoginReqBody(req)
	if err != nil {
		return nil, errors.ErrEncodingRequest
	}
	err = tcp.SendRequest(conn, location, header, encodedLoginRequest)
	if err != nil {
		return nil, errors.ErrWritingToServer
	}

	buffer, err := tcp_service.ReadResponseFromServer(conn)
	if err != nil {
		return nil, errors.ErrReadingResponse
	}
	response, err := tcp.DecodeTCPResponse(buffer)
	if err != nil {
		return nil, errors.ErrDecodingResponse
	}
	if response.StatusCode != tcp.StatusOK {
		return nil, tcp_service.ResponseErrorProduction(response.Data)
	}
	responseData, err := tcp.DecodeLoginResponse(response.Data)
	if err != nil {
		return nil, errors.ErrDecodingResponse
	}
	var token *models.Token
	token, err = tcp.DecodeToken(responseData.AuthToken)

	if err != nil {
		return nil, errors.ErrDecodingToken
	}
	return token, nil
}

func (s *APIService) AddCard(reqBody *tcp.AddCardBody) (*models.CreditCard, error) {
	location := "wallets/cards"
	header := make(map[string]string)
	methodHeader := tcp.MethodPost
	tcp_service.SetMethodHeader(header, methodHeader)

	conn, err := s.MakeNewTCPConnection()
	if err != nil {
		return nil, errors.ErrConnectionFailed
	}
	defer conn.Close()

	tcp_service.SetAuthorizationHeader(header)
	addCardReqBody := tcp.NewAddCardBody(reqBody.CardNumber)
	encodedAddCardReqBody, err := tcp.EncodeAddCardReqBody(addCardReqBody)
	if err != nil {
		return nil, errors.ErrEncodingRequest
	}
	err = tcp.SendRequest(conn, location, header, encodedAddCardReqBody)
	if err != nil {
		return nil, errors.ErrWritingToServer
	}

	buffer, err := tcp_service.ReadResponseFromServer(conn)
	if err != nil {
		return nil, errors.ErrReadingResponse
	}

	response, err := tcp.DecodeTCPResponse(buffer)
	if err != nil {
		return nil, errors.ErrDecodingResponse
	}
	if response.StatusCode != tcp.StatusCreated {
		return nil, tcp_service.ResponseErrorProduction(response.Data)
	}
	var addCardResBody *tcp.AddCardResponse
	addCardResBody, err = tcp.DecodeAddCardResponse(response.Data)
	if err != nil {
		return nil, errors.ErrDecodingSuccessfulResponse
	}
	newCard, err := tcp.DecodeCreditCard(addCardResBody.Card)
	if err != nil {
		return nil, errors.ErrDecodingSuccessfulResponse
	}
	return newCard, nil
}

func (s *APIService) Logout(req *tcp.LogoutUserReq) error {
	//TODO implement me
	return nil
}

func (s *APIService) DisplayCards() ([]*models.CreditCard, error) {
	location := "wallets/cards"
	header := make(map[string]string)
	methodHeader := tcp.MethodGet
	tcp_service.SetMethodHeader(header, methodHeader)

	conn, err := s.MakeNewTCPConnection()
	if err != nil {
		return nil, errors.ErrConnectionFailed
	}
	defer conn.Close()
	tcp_service.SetAuthorizationHeader(header)
	err = tcp.SendRequest(conn, location, header, nil)
	if err != nil {
		return nil, errors.ErrWritingToServer
	}

	// Read the response from the server
	buffer, err := tcp_service.ReadResponseFromServer(conn)
	if err != nil {
		return nil, errors.ErrReadingResponse
	}

	tcpResponse, err := tcp.DecodeTCPResponse(buffer)
	if err != nil {
		return nil, errors.ErrDecodingResponse
	}
	if tcpResponse.StatusCode != tcp.StatusOK {
		return nil, tcp_service.ResponseErrorProduction(tcpResponse.Data)
	}
	getCardsResBody, err := tcp.DecodeGetCardsBodyResponse(tcpResponse.Data)
	if err != nil {
		return nil, errors.ErrDecodingSuccessfulResponse
	}
	return getCardsResBody.Cards, nil
}

func (s *APIService) Deposit(data *tcp.DepositBody) error {
	location := "wallets/deposit"
	header := make(map[string]string)
	methodHeader := tcp.MethodPost
	tcp_service.SetMethodHeader(header, methodHeader)

	conn, err := s.MakeNewTCPConnection()
	if err != nil {
		return errors.ErrConnectionFailed
	}
	defer conn.Close()
	tcp_service.SetAuthorizationHeader(header)
	encodedWithdrawRepositReqBody, err := tcp.EncodeDepositReqBody(data)
	if err != nil {
		return errors.ErrEncodingRequest
	}
	err = tcp.SendRequest(conn, location, header, encodedWithdrawRepositReqBody)
	if err != nil {
		return errors.ErrWritingToServer
	}

	buffer, err := tcp_service.ReadResponseFromServer(conn)
	if err != nil {
		return errors.ErrReadingResponse
	}

	response, err := tcp.DecodeTCPResponse(buffer)
	if err != nil {
		return errors.ErrDecodingResponse
	}
	if response.StatusCode != tcp.StatusOK {
		return tcp_service.ResponseErrorProduction(response.Data)
	}
	return nil

}

func (s *APIService) Withdraw(data *tcp.WithdrawBody) error {
	location := "wallets/withdraw"
	header := make(map[string]string)
	methodHeader := tcp.MethodPost
	tcp_service.SetMethodHeader(header, methodHeader)

	conn, err := s.MakeNewTCPConnection()
	if err != nil {
		return errors.ErrConnectionFailed
	}
	defer conn.Close()
	tcp_service.SetAuthorizationHeader(header)
	encodedWithdrawRepositReqBody, err := tcp.EncodeWithdrawReqBody(data)
	if err != nil {
		return errors.ErrEncodingRequest
	}
	err = tcp.SendRequest(conn, location, header, encodedWithdrawRepositReqBody)
	if err != nil {
		return errors.ErrWritingToServer
	}

	buffer, err := tcp_service.ReadResponseFromServer(conn)
	if err != nil {
		return errors.ErrReadingResponse
	}

	response, err := tcp.DecodeTCPResponse(buffer)
	if err != nil {
		return errors.ErrDecodingResponse
	}
	if response.StatusCode != tcp.StatusOK {
		return tcp_service.ResponseErrorProduction(response.Data)
	}
	return nil
}

func (s *APIService) GetWallet() (*models.Wallet, error) {
	location := "wallets"
	header := make(map[string]string)
	methodHeader := tcp.MethodGet
	tcp_service.SetMethodHeader(header, methodHeader)

	conn, err := s.MakeNewTCPConnection()
	if err != nil {
		return nil, errors.ErrConnectionFailed
	}
	defer conn.Close()
	tcp_service.SetAuthorizationHeader(header)

	err = tcp.SendRequest(conn, location, header, nil)
	if err != nil {
		return nil, errors.ErrWritingToServer
	}

	// Read the response from the server
	buffer, err := tcp_service.ReadResponseFromServer(conn)
	if err != nil {
		return nil, errors.ErrReadingResponse
	}

	tcpResponse, err := tcp.DecodeTCPResponse(buffer)
	if err != nil {
		return nil, errors.ErrDecodingResponse
	}
	if tcpResponse.StatusCode != tcp.StatusOK {
		return nil, tcp_service.ResponseErrorProduction(tcpResponse.Data)
	}
	wallet, err := tcp.DecodeTCPWalletResponse(tcpResponse.Data)

	if err != nil {
		return nil, errors.ErrDecodingSuccessfulResponse
	}
	return wallet.Wallet, nil
}

func (s *APIService) CreateRestaurant(newRestaurant *models.Restaurant) (*models.Restaurant, error) {
	location := "restaurants"
	header := make(map[string]string)
	methodHeader := tcp.MethodPost
	tcp_service.SetMethodHeader(header, methodHeader)
	tcp_service.SetAuthorizationHeader(header)

	conn, err := s.MakeNewTCPConnection()
	if err != nil {
		return nil, errors.ErrConnectionFailed
	}
	defer conn.Close()
	// Send the message to the server
	createRestaurantRequestBody := tcp.CreateRestaurantBody{
		Name:    newRestaurant.Name,
		Phone:   newRestaurant.Phone,
		Address: newRestaurant.Address,
	}
	encodedCreateRestaurantRequestBody, err := tcp.EncodeCreateRestaurantRequest(createRestaurantRequestBody)
	if err != nil {
		return nil, errors.ErrEncodingRequest
	}
	err = tcp.SendRequest(conn, location, header, encodedCreateRestaurantRequestBody)
	if err != nil {
		return nil, errors.ErrWritingToServer
	}
	// Read the response from the server
	buffer, err := tcp_service.ReadResponseFromServer(conn)
	if err != nil {
		return nil, errors.ErrReadingResponse
	}

	response, err := tcp.DecodeTCPResponse(buffer)
	if err != nil {
		return nil, errors.ErrDecodingResponse
	}
	if response.StatusCode != tcp.StatusCreated {
		return nil, tcp_service.ResponseErrorProduction(response.Data)
	}

	var responseData tcp.CreateRestaurantResponse
	responseData, err = tcp.DecodeCreateRestaurantResponse(response.Data)
	if err != nil {
		return nil, errors.ErrDecodingSuccessfulResponse
	}
	return responseData.Restaurant, nil
}
