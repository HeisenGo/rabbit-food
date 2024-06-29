package commands

import (
	"client/models"
	"client/services"
)

type GetRestaurantsIHaveARoleCommand struct {
	service services.Service
}

func (c *GetRestaurantsIHaveARoleCommand) Execute() ([]*models.Restaurant, error) {

	restaurants, err := c.service.GetRestaurantsIHaveRoleIn()
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func NewGetRestaurantsIHaveARoleCommand(service services.Service) *GetRestaurantsIHaveARoleCommand {
	return &GetRestaurantsIHaveARoleCommand{service: service}
}
