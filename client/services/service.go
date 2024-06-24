package services

import (
	"client/models"
	"client/protocol/tcp"
)

type Service interface {
	Register(*models.User) (tcp.Token, error)
	Login(req *models.LoginUserReq) (*tcp.Token, error)
	Logout(req *models.LogoutUserReq) error
	GetWallet(req *models.GetWalletReq) (*models.Wallet, error)
}
