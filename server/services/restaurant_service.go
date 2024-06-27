package services

import (
	"context"
	"server/internal/models/restaurant/menu"
	"server/internal/models/restaurant/restaurant"
)

type RestaurantService struct {

	restaurantOps *restaurant.Ops
	menuOps       *menu.Ops
}

func NewRestaurantService(restaurantOps *restaurant.Ops, menuOps *menu.Ops) *RestaurantService {
	return &RestaurantService{
		restaurantOps: restaurantOps,
		menuOps:       menuOps,
	}
}

func (s *RestaurantService) CreateRestaurantForOwner(ctx context.Context, restaurant *restaurant.Restaurant) (*restaurant.Restaurant, error) {
	createdRestaurant, err := s.restaurantOps.Create(ctx, restaurant)
	if err != nil {
		return nil, err
	}
	return createdRestaurant, nil
}

func (s *RestaurantService) CreateMenuForRestaurant(ctx context.Context, menu *menu.Menu) (*menu.Menu, error) {
	createdMenu, err := s.menuOps.Create(ctx, menu)
	if err != nil {
		return nil, err
	}
	return createdMenu, nil
}

func (s *RestaurantService) GetAllRestaurantMenus(ctx context.Context, restaurant *restaurant.Restaurant) ([]*menu.Menu, error) {
	fetchedMenus, err := s.menuOps.GetAllRestaurantMenus(ctx, restaurant)
	if err != nil {
		return nil, err
	}
	return fetchedMenus, nil
}

func (s *RestaurantService) AddMenuItemToMenu(ctx context.Context, menuItem *menu.MenuItem) (*menu.MenuItem, error) {
	addedMenuItem, err := s.menuOps.AddMenuItemToMenu(ctx, menuItem)
	if err != nil {
		return nil, err
	}
	return addedMenuItem, nil
}

func (s *RestaurantService) GetMenuItemsOfMenu(ctx context.Context, menu *menu.Menu) ([]*menu.MenuItem, error) {
	fetchedMenuItems, err := s.menuOps.GetMenuItemsOfMenu(ctx, menu)
	if err != nil {
		return nil, err
	}
	return fetchedMenuItems, nil
}
