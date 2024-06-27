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
