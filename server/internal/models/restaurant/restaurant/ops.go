package restaurant

import (
	"context"

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

func (o *Ops) Create(ctx context.Context, restaurant *Restaurant) (*Restaurant, error) {
	return o.repo.CreateRestaurantAndAssignOwner(ctx, restaurant)
}

func (o *Ops) AddCategoriesToRestaurant(ctx context.Context, rest *Restaurant, categoryIDs []uint) (*Restaurant, error) {
	return o.repo.AddCategoriesToRestaurant(ctx, rest, categoryIDs)
}

func (o *Ops) GetRestaurantCategories(ctx context.Context, restaurantID uint) ([]*RestaurantCategory, error) {
	return o.repo.GetRestaurantCategories(ctx, restaurantID)
}
