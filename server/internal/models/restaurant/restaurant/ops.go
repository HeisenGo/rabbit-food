package restaurant

import (
	"context"
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
