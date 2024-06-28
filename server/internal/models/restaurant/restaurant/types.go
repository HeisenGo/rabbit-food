package restaurant

import (
	"context"
	"server/internal/models/address"
)

type Restaurant struct {
	ID      uint
	Name    string
	Phone   string
	Address *address.Address
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
