package services

import (
	"context"
	"server/internal/models/restaurant/restaurant"
	"server/internal/models/user"
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
	createdRestaurant, err := s.restauarntOps.Create(ctx, restaurant)
	if err != nil {
		return nil, err
	}
	return createdRestaurant, nil
}

func (s *RestaurantService) IsRestaurantOwner(ctx context.Context, restaurantID uint) (bool, error) {
	isOk, err := s.restauarntOps.IsRestaurantOwner(ctx, restaurantID)
	if err != nil {
		return false, err
	}
	return isOk, nil
}

func (s *RestaurantService) GetRestaurantByID(ctx context.Context, restaurantID uint) (*restaurant.Restaurant, error) {
	restaurant, err := s.restauarntOps.GetByID(ctx, restaurantID)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}

func (s *RestaurantService) AssignOperatorToRestarant(ctx context.Context, operator *user.User, restaurant restaurant.Restaurant) (*user.User, error) {
	user, err := s.restauarntOps.AssignOperatorToRestarant(ctx, operator, restaurant)
	if err != nil {
		return nil, err
	}
	return user, nil
}
