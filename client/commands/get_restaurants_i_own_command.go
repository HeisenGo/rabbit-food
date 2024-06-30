package commands

import (
	"client/models"
	"client/services"
)

type GetRestaurantsIOwnCommand struct {
	service services.Service
}

func (c *GetRestaurantsIOwnCommand) Execute() ([]*models.Restaurant, error) {

	restaurants, err := c.service.GetRestaurantsIOwn()
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func NewGetRestaurantsIOwnCommand(service services.Service) *GetRestaurantsIOwnCommand {
	return &GetRestaurantsIOwnCommand{service: service}
}
