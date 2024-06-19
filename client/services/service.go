package services

import "client/models"

type Service interface {
	Register(*models.User) error
	GetWallet(req *models.GetWalletReq) (*models.Wallet, error)
}
