package commands

import (
	"client/protocol/tcp"
	"client/services"
)

type AddOperatorToRestaurantCommand struct {
	service services.Service
}

func (c *AddOperatorToRestaurantCommand) Execute(addOperatorData tcp.RestaurantAddOperatorReqBody) error {

	err := c.service.AddOperatorToRestaurant(&addOperatorData)
	if err != nil {
		return err
	}
	return nil
}

func NewAddOperatorToRestaurantCommand(service services.Service) *AddOperatorToRestaurantCommand {
	return &AddOperatorToRestaurantCommand{service: service}
}
