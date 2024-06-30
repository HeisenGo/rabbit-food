package commands

import (
	"client/protocol/tcp"
	"client/services"
)

type AddMotorToRestaurantCommand struct {
	service services.Service
}

func (c *AddMotorToRestaurantCommand) Execute(addMotorData tcp.RestaurantMotorReqBody) error {

	err := c.service.AddMotorToRestaurant(&addMotorData)
	if err != nil {
		return err
	}
	return nil
}

func NewAddMotorToRestaurantCommand(service services.Service) *AddMotorToRestaurantCommand {
	return &AddMotorToRestaurantCommand{service: service}
}
