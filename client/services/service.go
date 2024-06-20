package services

import "client/models"

type Service interface {
	Register(*models.User) error
	Login(req *models.LoginUserReq) (*models.User, error)
	Logout(req *models.LogoutUserReq) error
	GetWallet(req *models.GetWalletReq) (*models.Wallet, error)
}
