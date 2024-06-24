package commands

import (
	"client/models"
	"client/services"
	"errors"
)

type GetWalletCommand struct {
	service services.Service
}

func (c *GetWalletCommand) Execute(walletData any) (*models.Wallet, error) {
	walletReq, ok := walletData.(*models.GetWalletReq)
	if !ok {
		return nil, errors.New("data type isn't GetWalletReq")
	}
	receivedWallet, err := c.service.GetWallet(walletReq)
	if err != nil {
		return nil, err
	}
	return receivedWallet, nil
}

func NewGetWalletCommand(service services.Service) *GetWalletCommand {
	return &GetWalletCommand{service: service}
}
