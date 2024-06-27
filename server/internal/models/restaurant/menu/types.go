package menu

import (
	"context"
	"server/internal/models/restaurant/restaurant"
)

type Menu struct {
	ID           uint
	Name         string
	RestaurantID uint
}

type MenuItem struct {
	ID                            uint
	Name                          string
	Price                         uint
	PreparationTime               uint
	CancellationPenaltyPercentage uint
	MenuID                        uint
}

func NewMenuItem(name string, price uint, preparationTime uint, cancellationPenaltyPercentage uint, menuID uint) *MenuItem {
	return &MenuItem{Name: name, Price: price, PreparationTime: preparationTime, CancellationPenaltyPercentage: cancellationPenaltyPercentage, MenuID: menuID}
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
