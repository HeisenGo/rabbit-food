package commands

import (
	"client/models"
	"client/services"
)

type GetMenusOfRestaurantCommand struct {
	service services.Service
}

func (c *GetMenusOfRestaurantCommand) Execute(restaurantID uint) ([]*models.RestaurantMenu, error) {

	menus, err := c.service.GetMenusOfRestaurant(restaurantID)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func NewGetMenusOfRestaurantCommand(service services.Service) *GetMenusOfRestaurantCommand {
	return &GetMenusOfRestaurantCommand{service: service}
}
