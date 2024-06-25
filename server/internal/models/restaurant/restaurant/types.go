package restaurant

import (
	"context"
	"server/internal/models/user"
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
	CheckMatchedRestaurantsOwnerIdAndClaimedID(ctx context.Context, restaurantID uint) (bool, error)
	GetByID(ctx context.Context, restaurantID uint) (*Restaurant, error)
	AssignOperatorToRestarant(ctx context.Context, operator *user.User, restaurant Restaurant) (*user.User, error)
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
