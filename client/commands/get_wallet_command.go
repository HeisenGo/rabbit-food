package commands

import (
	"client/models"
	"client/services"
	"fmt"
)

type GetWalletCommand struct {
	service services.Service
}

func (c *GetWalletCommand) Execute() (*models.Wallet, error) {

	fmt.Println("\nBallance is:")
	wallet, err := c.service.GetWallet()
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func NewGetWalletCommand(service services.Service) *GetWalletCommand {
	return &GetWalletCommand{service: service}
}
