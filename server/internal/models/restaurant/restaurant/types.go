package restaurant

import (
	"context"
)

type Restaurant struct {
	ID          uint
	Name        string
	Phone       string
	City        string
	Address     string
	Coordinates string // To do
	//RelatedUserID uint
	//UserRole      server.RoleType
	//UserID uint
	//Users []user.User
}
type Repo interface {
	CreateRestaurantAndAssignOwner(ctx context.Context, restauran *Restaurant) (*Restaurant, error)
}

func NewRestaurant(name string, phone string, city string, address string, coordinates string) *Restaurant {
	restaurant := &Restaurant{
		Name:        name,
		Phone:       phone,
		City:        city,
		Address:     address,
		Coordinates: coordinates,
	}
	return restaurant
}

func NewRestaurantByID(id uint) *Restaurant {
	restaurant := &Restaurant{
		ID: id,
	}
	return restaurant
}
