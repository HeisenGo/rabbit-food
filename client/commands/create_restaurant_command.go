package commands

import (
	"client/models"
	"client/services"
)

type CreateRestaurantCommand struct {
	service services.Service
}

func (c *CreateRestaurantCommand) Execute(createRestaurantData models.Restaurant) (*models.Restaurant, error) {

	newRestaurant, err := c.service.CreateRestaurant(&createRestaurantData)
	if err != nil {
		return nil, err
	}
	return newRestaurant, nil
}

func NewCreateRestaurantCommand(service services.Service) *CreateRestaurantCommand {
	return &CreateRestaurantCommand{service: service}
}
