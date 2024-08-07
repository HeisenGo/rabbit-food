package restaurant

import (
	"context"
	"server/internal/models/address"
	"server/internal/models/restaurant/motor"
	"server/internal/models/user"
)

type Restaurant struct {
	ID      uint             `json:"id"`
	Name    string           `json:"name"`
	Phone   string           `json:"phone"`
	Address *address.Address `json:"address"`
}

type RestaurantCategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Repo interface {
	CreateRestaurantAndAssignOwner(ctx context.Context, restauran *Restaurant) (*Restaurant, error)
	CheckMatchedRestaurantsOwnerIdAndClaimedID(ctx context.Context, restaurantID uint) (bool, error)
	GetByID(ctx context.Context, restaurantID uint) (*Restaurant, error)
	AssignOperatorToRestaurant(ctx context.Context, operator *user.User, restaurant Restaurant) (*user.User, error)
	RemoveOperatorFromRestaurant(ctx context.Context, operatorID uint, restaurantID uint) error
	WithdrawRestaurant(ctx context.Context, newOwnerID uint, restaurantID uint) error
	AddMotor(ctx context.Context, motor *motor.Motor, restaurantID uint) (*motor.Motor, error)
	RemoveMotor(ctx context.Context, motorID uint) error
	GetAllMotors(ctx context.Context, restaurantID uint) ([]*motor.Motor, error)
	GetAllOperators(ctx context.Context, restaurantID uint) ([]*user.User, error)
	DoesThisHaveARoleInRestaurant(ctx context.Context, restaurantID uint) (bool, error)
	GetOwnerInfo(ctx context.Context, restaurantID uint) (*user.User, error)
	GetRestaurantInfo(ctx context.Context, restaurantID uint) (*Restaurant, *user.User, []*user.User, []*motor.Motor, error)
	RemoveRestaurant(ctx context.Context, restaurantID uint) error
	GetRestaurantsOfAnOwner(ctx context.Context) ([]*Restaurant, error)
	GetRestaurantsOfAnOperator(ctx context.Context) ([]*Restaurant, error)
	EditRestaurantName(ctx context.Context, restaurantID uint, newName string) error
	GetRestaurantsToAddCategoryMenuFood(ctx context.Context) ([]*Restaurant, error)
  AddCategoriesToRestaurant(ctx context.Context, rest *Restaurant, categoryIDs []uint) (*Restaurant, error)
	GetRestaurantCategories(ctx context.Context, restaurantID uint) ([]*RestaurantCategory, error)
}

func NewRestaurant(name string, phone string, address address.Address) *Restaurant {
	restaurant := &Restaurant{
		Name:    name,
		Phone:   phone,
		Address: &address,
	}
	return restaurant
}

func NewRestaurantByID(id uint) *Restaurant {
	restaurant := &Restaurant{
		ID: id,
	}
	return restaurant
}
