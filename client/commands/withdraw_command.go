package commands

import (
	"client/protocol/tcp"
	"client/services"
)

type WithdrawCommand struct {
	service services.Service
}

func (c *WithdrawCommand) Execute(data *tcp.WithdrawBody) error {

	err := c.service.Withdraw(data)
	if err != nil {
		return err
	}
	return nil
}

func NewWithdrawCommand(service services.Service) *WithdrawCommand {
	return &WithdrawCommand{service: service}
}
