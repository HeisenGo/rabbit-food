package restaurant

import (
	"context"
	"server/internal/models/restaurant/motor"
	"server/internal/models/user"

	"gorm.io/gorm"
)

type Ops struct {
	db   *gorm.DB
	repo Repo
}

func NewRestaurantOps(db *gorm.DB, repo Repo) *Ops {
	return &Ops{
		db:   db,
		repo: repo,
	}
}

func (o *Ops) Create(ctx context.Context, restauarnt *Restaurant) (*Restaurant, error) {
	return o.repo.CreateRestaurantAndAssignOwner(ctx, restauarnt)
}

func (o *Ops) IsRestaurantOwner(ctx context.Context, restaurantID uint) (bool, error) {
	return o.repo.CheckMatchedRestaurantsOwnerIdAndClaimedID(ctx, restaurantID)
}

func (o *Ops) GetByID(ctx context.Context, restaurantID uint) (*Restaurant, error) {
	return o.repo.GetByID(ctx, restaurantID)
}

func (o *Ops) AssignOperatorToRestarant(ctx context.Context, operator *user.User, restaurant Restaurant) (*user.User, error) {
	return o.repo.AssignOperatorToRestarant(ctx, operator, restaurant)
}

func (o *Ops) RemoveOperatorFromRestarant(ctx context.Context, operatorID uint, restaurantID uint) error {
	return o.repo.RemoveOperatorFromRestarant(ctx, operatorID, restaurantID)
}

func (o *Ops) AddMotor(ctx context.Context, motor *motor.Motor, restaurantID uint) (*motor.Motor, error) {
	return o.repo.AddMotor(ctx, motor, restaurantID)
}

func (o *Ops) RemoveMotor(ctx context.Context, motorID uint) error {
	return o.repo.RemoveMotor(ctx, motorID)
}

func (o *Ops) WithdrawRestaurant(ctx context.Context, newOwnerID uint, restaurantID uint) error {
	return o.repo.WithdrawRestaurant(ctx, newOwnerID, restaurantID)
}

func (o *Ops) GetAllMotors(ctx context.Context, restaurantID uint) ([]*motor.Motor, error) {
	return o.repo.GetAllMotors(ctx, restaurantID)
}
func (o *Ops) GetAllOperators(ctx context.Context, restaurantID uint) ([]*user.User, error) {
	return o.repo.GetAllOperators(ctx, restaurantID)
}
func (o *Ops) DoeseThisHaveARoleInRestaurant(ctx context.Context, restaurantID uint) (bool, error) {
	return o.repo.DoeseThisHaveARoleInRestaurant(ctx, restaurantID)
}
func (o *Ops) GetOwnerInfo(ctx context.Context, restaurantID uint) (*user.User, error) {
	return o.repo.GetOwnerInfo(ctx, restaurantID)
}
func (o *Ops) GetRestarantInfo(ctx context.Context, restaurantID uint) (*Restaurant,
	*user.User, []*user.User, []*motor.Motor, error) {
	return o.repo.GetRestarantInfo(ctx, restaurantID)
}

func (o *Ops) RemoveRestaurant(ctx context.Context, restaurantID uint) error {
	return o.repo.RemoveRestaurant(ctx, restaurantID)
}

func (o *Ops) GetRestaurantsOfAnOwner(ctx context.Context) ([]*Restaurant, error){
	return o.repo.GetRestaurantsOfAnOwner(ctx)
}
func (o *Ops) GetRestaurantsOfAnOperator(ctx context.Context) ([]*Restaurant, error){
	return o.repo.GetRestaurantsOfAnOperator(ctx)
}

func (o *Ops) EditRestaurantName(ctx context.Context, restaurantID uint, newName string) error{

	return o.repo.EditRestaurantName(ctx, restaurantID, newName)
}
