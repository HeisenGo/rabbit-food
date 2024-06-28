package restaurant

import (
	"context"
	"server/internal/models/address"
)

type Restaurant struct {
	ID         uint                  `json:"id"`
	Name       string                `json:"name"`
	Phone      string                `json:"phone"`
	City       string                `json:"city"`
	Address    address.Address       `json:"address"`
	Categories []*RestaurantCategory `json:"categories"`
}

type RestaurantCategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Repo interface {
	CreateRestaurantAndAssignOwner(ctx context.Context, restaurant *Restaurant) (*Restaurant, error)
	AddCategoriesToRestaurant(ctx context.Context, rest *Restaurant, categoryIDs []uint) (*Restaurant, error)
	GetRestaurantCategories(ctx context.Context, restaurantID uint) ([]*RestaurantCategory, error)
}

func NewRestaurant(name string, phone string, city string, address address.Address) *Restaurant {
	restaurant := &Restaurant{
		Name:    name,
		Phone:   phone,
		City:    city,
		Address: address,
	}
	return restaurant
}

func NewRestaurantByID(id uint) *Restaurant {
	restaurant := &Restaurant{
		ID: id,
	}
	return restaurant
}
