package menu

import (
	"context"
	"server/internal/models/restaurant/restaurant"
)

type Menu struct {
	ID           uint   `json:"menu_id"`
	Name         string `json:"name"`
	RestaurantID uint   `json:"restaurant_id"`
}

type MenuItem struct {
	ID                            uint   `json:"id"`
	Name                          string `json:"name"`
	Price                         uint   `json:"price"`
	PreparationMinutes            uint   `json:"preparation_minutes"` // in minutes
	CancellationPenaltyPercentage uint   `json:"cancellation_penalty_percentage"`
	MenuID                        uint   `json:"menu_id"`
}

func NewMenuItem(name string, price uint, preparationMinutes uint, cancellationPenaltyPercentage uint, menuID uint) *MenuItem {
	return &MenuItem{Name: name, Price: price, PreparationMinutes: preparationMinutes, CancellationPenaltyPercentage: cancellationPenaltyPercentage, MenuID: menuID}
}

type Repo interface {
	CreateMenu(ctx context.Context, menu *Menu) (*Menu, error)
	GetAllRestaurantMenus(ctx context.Context, restaurant *restaurant.Restaurant) ([]*Menu, error)
	AddMenuItemToMenu(ctx context.Context, menuItem *MenuItem) (*MenuItem, error)
	GetMenuItemsOfMenu(ctx context.Context, menu *Menu) ([]*MenuItem, error)
}

func NewMenu(name string, restaurantID uint) *Menu {
	menu := &Menu{
		Name:         name,
		RestaurantID: restaurantID,
	}
	return menu
}

func NewMenuByID(id uint) *Menu {
	menu := &Menu{
		ID: id,
	}
	return menu
}
