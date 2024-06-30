package commands

import (
	"client/models"
	"client/services"
)

type GetItemsOfMenuCommand struct {
	service services.Service
}

func (c *GetItemsOfMenuCommand) Execute(menuID uint) ([]*models.MenuItem, error) {

	items, err := c.service.GetItemsOfMenu(menuID)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func NewGetItemsOfMenuCommand(service services.Service) *GetMenusOfRestaurantCommand {
	return &GetMenusOfRestaurantCommand{service: service}
}
