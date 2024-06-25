package restaurant

import (
	"context"
	"server"
)

type Repo interface {
	Create(ctx context.Context, userRest *UserRestaurant) (*UserRestaurant, error)
}

type UserRestaurant struct {
	ID           uint
	UserID       uint
	RestaurantID uint
	RoleType     string
}

func NewUserRestaurant(userID uint, restaurantID uint, userRole server.RoleType) *UserRestaurant {
	return &UserRestaurant{UserID: userID, RestaurantID: restaurantID, RoleType: string(userRole)}
}
