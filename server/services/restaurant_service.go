package services

import (
	"context"
	"server/internal/models/restaurant/restaurant"
)

type RestaurantService struct {
	restauarntOps *restaurant.Ops
}

func NewRestaurantService(restauarntOps *restaurant.Ops) *RestaurantService {
	return &RestaurantService{
		restauarntOps: restauarntOps,
	}
}

func (s *RestaurantService) CreateResturantForOwner(ctx context.Context, restaurant *restaurant.Restaurant) (*restaurant.Restaurant, error) {
	createdWallet, err := s.restauarntOps.Create(ctx, restaurant)
	if err != nil {
		return nil, err
	}
	return createdWallet, nil
}
