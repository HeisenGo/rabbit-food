package commands

import (
	"client/protocol/tcp"
	"client/services"
)

type AddCategoryToRestaurantCommand struct {
	service services.Service
}

func (c *AddCategoryToRestaurantCommand) Execute(addCategoryData tcp.RestaurantCategoryBody) error {

	err := c.service.AddCategoryToRestaurant(&addCategoryData)
	if err != nil {
		return err
	}
	return nil
}

func NewAddCategoryToRestaurantCommand(service services.Service) *AddCategoryToRestaurantCommand {
	return &AddCategoryToRestaurantCommand{service: service}
}
