package restaurant

import (
	"context"
	"server/internal/models/restaurant/motor"
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
	RemoveOperatorFromRestarant(ctx context.Context, operatorID uint, restaurantID uint) error
	WithdrawRestaurant(ctx context.Context, newOwnerID uint, restaurantID uint) error
	AddMotor(ctx context.Context, motor *motor.Motor, restaurantID uint) (*motor.Motor, error)
	RemoveMotor(ctx context.Context, motorID uint) error
	GetAllMotors(ctx context.Context, restaurantID uint) ([]*motor.Motor, error)
	GetAllOperators(ctx context.Context, restaurantID uint) ([]*user.User, error)
	DoeseThisHaveARoleInRestaurant(ctx context.Context, restaurantID uint) (bool, error)
	GetOwnerInfo(ctx context.Context, restaurantID uint) (*user.User, error) 
	GetRestarantInfo(ctx context.Context, restaurantID uint) (*Restaurant, *user.User, []*user.User, []*motor.Motor, error) 
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
