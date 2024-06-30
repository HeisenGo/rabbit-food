package commands

import (
	"client/models"
	"client/services"
)

type GetCategoriesOfRestaurantCommand struct {
	service services.Service
}

func (c *GetCategoriesOfRestaurantCommand) Execute(restaurantID uint) ([]*models.RestaurantCategory, error) {

	cats, err := c.service.GetCategoriesOfRestaurant(restaurantID)
	if err != nil {
		return nil, err
	}
	return cats, nil
}

func NewGetCategoriesOfRestaurantCommand(service services.Service) *GetCategoriesOfRestaurantCommand {
	return &GetCategoriesOfRestaurantCommand{service: service}
}
