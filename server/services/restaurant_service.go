package services

import (
	"context"
	"server/internal/models/restaurant"
)

type RestaurantService struct {
	restauarntOps *restaurant.Ops
}

func NewRestaurantService(restauarntOps *restaurant.Ops) *RestaurantService {
	return &RestaurantService{
		restauarntOps: restauarntOps,
	}
}

func (s *RestaurantService) CreateResturantByUserID(ctx context.Context, name string, userID uint, roleType string) (*restaurant.Restaurant, error) {
	newRestaurant := restaurant.NewRestaurant(name, userID, roleType)
	createdWallet, err := s.restauarntOps.Create(ctx, newRestaurant)
	if err != nil {
		return nil, err
	}
	return createdWallet, nil
}
