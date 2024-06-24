package restaurant

import (
	"context"
)

type Restaurant struct {
	ID   uint
	Name string
	//UserID uint
	//Users []user.User
}

type UserRestaurant struct {
	ID           uint
	UserID       uint
	RestaurantID uint
	RoleType     string
}

type Repo interface {
	Create(ctx context.Context, restaurant *Restaurant) (*Restaurant, error)
}

func NewRestaurant(name string, userID uint, roleType string) *Restaurant {
	//users := []User
	//users = append(users, user)
	restaurant := &Restaurant{
		Name: name,
		//Users : users,
	}
	return restaurant
}
