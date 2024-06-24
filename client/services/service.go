package services

import (
	"client/models"
	"client/protocol/tcp"
)

type Service interface {
	Register(*models.User) (*models.Token, error)
	Login(req *tcp.LoginBody) (*models.Token, error)
	Logout(req *tcp.LogoutUserReq) error
	GetWallet(req *models.GetWalletReq) (*models.Wallet, error)
	AddCard(req *tcp.AddCardBody) (*models.CreditCard, error)
}
