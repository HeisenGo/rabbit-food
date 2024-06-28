package restaurant

import (
	"context"
	"server/internal/models/address"
)

type Restaurant struct {
	ID      uint             `json:"id"`
	Name    string           `json:"name"`
	Phone   string           `json:"phone"`
	Address *address.Address `json:"address"`
}
type Repo interface {
	CreateRestaurantAndAssignOwner(ctx context.Context, restauran *Restaurant) (*Restaurant, error)
}

func NewRestaurant(name string, phone string, address address.Address) *Restaurant {
	restaurant := &Restaurant{
		Name:    name,
		Phone:   phone,
		Address: &address,
	}
	return restaurant
}

func NewRestaurantByID(id uint) *Restaurant {
	restaurant := &Restaurant{
		ID: id,
	}
	return restaurant
}
