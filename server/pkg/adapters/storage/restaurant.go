package storage

import (
	"context"
	"errors"
	"server/internal/errors/users"
	"server/internal/models/restaurant"
	"server/pkg/adapters/storage/mappers"

	"gorm.io/gorm"
)

type restaurantRepo struct {
	db *gorm.DB
}

func NewRestaurantRepo(db *gorm.DB) restaurant.Repo {
	return &restaurantRepo{
		db: db,
	}
}

func (r *restaurantRepo) Create(ctx context.Context, restaurant *restaurant.Restaurant) (*restaurant.Restaurant, error) {
	newRestaurant := mappers.RestaurantDomainToEntity(restaurant)
	err := r.db.Create(&newRestaurant).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, users.ErrUserExists
		}
		return nil, err
	}
	createdRestaurant := mappers.RestaurantEntityToDomain(newRestaurant)
	return createdRestaurant, nil
}
