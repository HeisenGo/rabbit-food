package services

import (
	"client/models"
	"client/protocol/tcp"
)

type Service interface {
	Register(*models.User) (*models.Token, error)
	Login(resq *tcp.LoginBody) (*models.Token, error)
	Logout(req *tcp.LogoutUserReq) error
	GetWallet() (*models.Wallet, error)
	AddCard(req *tcp.AddCardBody) (*models.CreditCard, error)
	DisplayCards() ([]*models.CreditCard, error)
	Deposit(data *tcp.DepositBody) error
	Withdraw(data *tcp.WithdrawBody) error
  DisplayProfile(userID uint) (*models.User, error)
	EditProfile(user *models.User) (*models.User, error)
}
