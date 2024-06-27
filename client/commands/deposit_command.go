package commands

import (
	"client/protocol/tcp"
	"client/services"
)

type DepositCommand struct {
	service services.Service
}

func (c *DepositCommand) Execute(data *tcp.DepositBody) error {

	err := c.service.Deposit(data)
	if err != nil {
		return err
	}
	return nil
}

func NewDepositCommand(service services.Service) *DepositCommand {
	return &DepositCommand{service: service}
}
