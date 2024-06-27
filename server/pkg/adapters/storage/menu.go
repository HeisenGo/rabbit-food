package storage

import (
	"context"
	"gorm.io/gorm"
	"server/internal/models/restaurant/menu"
	"server/internal/models/restaurant/restaurant"
	"server/pkg/adapters/storage/entities"
	"server/pkg/adapters/storage/mappers"
)

type manuRepo struct {
	db *gorm.DB
}

func NewMenuRepo(db *gorm.DB) menu.Repo {
	return &manuRepo{
		db: db,
	}
}

func (r *manuRepo) CreateMenu(ctx context.Context, menu *menu.Menu) (*menu.Menu, error) {
	newMenu := mappers.MenuDomainToEntity(menu)
	// TODO: check restaurant permission
	err := r.db.Create(&newMenu).Error
	if err != nil {
		return nil, err
	}
	createdMenu := mappers.MenuEntityToDomain(newMenu)
	return createdMenu, nil
}

func (r *manuRepo) GetAllRestaurantMenus(ctx context.Context, restaurant *restaurant.Restaurant) ([]*menu.Menu, error) {
	var menuEntities []*entities.Menu
	err := r.db.WithContext(ctx).Where("restaurant_id = ?", restaurant.ID).Find(&menuEntities).Error
	if err != nil {
		return nil, err
	}
	domainMenus := mappers.BatchMenuEntityToDomain(menuEntities)
	return domainMenus, nil
}

func (r *manuRepo) AddMenuItemToMenu(ctx context.Context, menuItem *menu.MenuItem) (*menu.MenuItem, error) {
	newMenuItem := mappers.MenuItemDomainToEntity(menuItem)
	// TODO: check restaurant permission
	err := r.db.Create(&newMenuItem).Error
	if err != nil {
		return nil, err
	}
	createdMenuItem := mappers.MenuItemEntityToDomain(newMenuItem)
	return createdMenuItem, nil
}
