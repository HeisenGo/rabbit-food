package services

import (
	"client/models"
	"client/protocol/tcp"
)

type Service interface {
	Register(*models.User) (*models.Token, error)
	Login(request *tcp.LoginBody) (*models.Token, error)
	Logout(req *tcp.LogoutUserReq) error
	GetWallet() (*models.Wallet, error)
	AddCard(req *tcp.AddCardBody) (*models.CreditCard, error)
	DisplayCards() ([]*models.CreditCard, error)
	Deposit(data *tcp.DepositBody) error
	Withdraw(data *tcp.WithdrawBody) error
	CreateRestaurant(newRestaurant *models.Restaurant) (*models.Restaurant, error)
	GetRestaurantsIHaveRoleIn() ([]*models.Restaurant, error)
}
