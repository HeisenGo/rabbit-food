package services

import (
	"context"
	"server/internal/models/restaurant/menu"
	"server/internal/models/restaurant/motor"
	"server/internal/models/restaurant/restaurant"
	"server/internal/models/user"
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

func (s *RestaurantService) IsRestaurantOwner(ctx context.Context, restaurantID uint) (bool, error) {
	isOk, err := s.restaurantOps.IsRestaurantOwner(ctx, restaurantID)
	if err != nil {
		return false, err
	}
	return isOk, nil
}

func (s *RestaurantService) GetRestaurantByID(ctx context.Context, restaurantID uint) (*restaurant.Restaurant, error) {
	restaurant, err := s.restaurantOps.GetByID(ctx, restaurantID)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}

func (s *RestaurantService) AssignOperatorToRestaurant(ctx context.Context, operator *user.User, restaurant restaurant.Restaurant) (*user.User, error) {
	user, err := s.restaurantOps.AssignOperatorToRestaurant(ctx, operator, restaurant)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *RestaurantService) RemoveOperatorFromRestaurant(ctx context.Context, operatorID uint, restaurantID uint) error {
	err := s.restaurantOps.RemoveOperatorFromRestaurant(ctx, operatorID, restaurantID)
	if err != nil {
		return err
	}
	return nil
}

func (s *RestaurantService) AddMotor(ctx context.Context, motor *motor.Motor, restaurantID uint) (*motor.Motor, error) {
	motor, err := s.restaurantOps.AddMotor(ctx, motor, restaurantID)
	if err != nil {
		return nil, err
	}
	return motor, nil
}

func (s *RestaurantService) RemoveMotor(ctx context.Context, motorID uint) error {
	err := s.restaurantOps.RemoveMotor(ctx, motorID)
	if err != nil {
		return err
	}
	return nil
}

func (s *RestaurantService) WithdrawRestaurant(ctx context.Context, newOwnerID uint, restaurantID uint) error {
	err := s.restaurantOps.WithdrawRestaurant(ctx, newOwnerID, restaurantID)
	if err != nil {
		return err
	}
	return nil
}

func (s *RestaurantService) GetAllMotors(ctx context.Context, restaurantID uint) ([]*motor.Motor, error) {
	motors, err := s.restaurantOps.GetAllMotors(ctx, restaurantID)
	if err != nil {
		return nil, err
	}
	return motors, nil
}

func (s *RestaurantService) GetAllOperators(ctx context.Context, restaurantID uint) ([]*user.User, error) {
	operators, err := s.restaurantOps.GetAllOperators(ctx, restaurantID)
	if err != nil {
		return nil, err
	}
	return operators, nil
}

func (s *RestaurantService) DoesThisHaveARoleInRestaurant(ctx context.Context, restaurantID uint) (bool, error) {
	yes, err := s.restaurantOps.DoesThisHaveARoleInRestaurant(ctx, restaurantID)
	if err != nil {
		return false, err
	}
	return yes, nil
}

func (s *RestaurantService) GetOwnerInfo(ctx context.Context, restaurantID uint) (*user.User, error) {
	owner, err := s.restaurantOps.GetOwnerInfo(ctx, restaurantID)
	if err != nil {
		return nil, err
	}
	return owner, nil
}

func (s *RestaurantService) GetRestaurantInfo(ctx context.Context, restaurantID uint) (*restaurant.Restaurant,
	*user.User, []*user.User, []*motor.Motor, error) {
	return s.restaurantOps.GetRestaurantInfo(ctx, restaurantID)
}

func (s *RestaurantService) RemoveRestaurant(ctx context.Context, restaurantID uint) error {
	return s.restaurantOps.RemoveRestaurant(ctx, restaurantID)
}

func (s *RestaurantService) GetRestaurantsOfAnOwner(ctx context.Context) ([]*restaurant.Restaurant, error) {
	return s.restaurantOps.GetRestaurantsOfAnOwner(ctx)
}
func (s *RestaurantService) GetRestaurantsOfAnOperator(ctx context.Context) ([]*restaurant.Restaurant, error) {
	return s.restaurantOps.GetRestaurantsOfAnOperator(ctx)
}

func (s *RestaurantService) EditRestaurantName(ctx context.Context, restaurantID uint, newName string) error {
	return s.restaurantOps.EditRestaurantName(ctx, restaurantID, newName)
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
func (s *RestaurantService)GetRestaurantsToAddCategoryMenuFood(ctx context.Context) ([]*restaurant.Restaurant, error){
	return s.restaurantOps.GetRestaurantsToAddCategoryMenuFood(ctx)
}