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

func NewRestaurant(name string, phone string, city string, address string, coordiantes string) *Restaurant {
	//users := []User
	//users = append(users, user)
	restaurant := &Restaurant{
		Name:        name,
		Phone:       phone,
		City:        city,
		Address:     address,
		Coordinates: coordiantes,
	}
	return restaurant
}
