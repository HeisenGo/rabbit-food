package menu

import (
	"context"
	"server/internal/models/restaurant/restaurant"

	"gorm.io/gorm"
)

type Ops struct {
	db   *gorm.DB
	repo Repo
}

func NewMenuOps(db *gorm.DB, repo Repo) *Ops {
	return &Ops{
		db:   db,
		repo: repo,
	}
}

func (o *Ops) Create(ctx context.Context, menu *Menu) (*Menu, error) {
	return o.repo.CreateMenu(ctx, menu)
}

func (o *Ops) GetAllRestaurantMenus(ctx context.Context, restaurant *restaurant.Restaurant) ([]*Menu, error) {
	return o.repo.GetAllRestaurantMenus(ctx, restaurant)
}

func (o *Ops) AddMenuItemToMenu(ctx context.Context, menuItem *MenuItem) (*MenuItem, error) {
	return o.repo.AddMenuItemToMenu(ctx, menuItem)
}
