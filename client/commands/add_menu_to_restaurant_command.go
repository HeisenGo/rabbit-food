package commands

import (
	"client/protocol/tcp"
	"client/services"
)

type AddMenuToRestaurantCommand struct {
	service services.Service
}

func (c *AddMenuToRestaurantCommand) Execute(addMenuData tcp.RestaurantMenuBody) error {

	err := c.service.AddMenuToRestaurant(&addMenuData)
	if err != nil {
		return err
	}
	return nil
}

func NewAddMenuToRestaurantCommand(service services.Service) *AddMenuToRestaurantCommand {
	return &AddMenuToRestaurantCommand{service: service}
}
