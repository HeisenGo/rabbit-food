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

// // NewRestaurant creates a new restaurant and associates it with an owner
// func (r *restaurantRepo) Create(ctx context.Context, restaurant *restaurant.Restaurant, ownerID uint) (*restaurant.Restaurant, error) {
// 	newRestaurant := mappers.RestaurantDomainToEntity(restaurant)
// 	// Begin a transaction
// 	tx := r.db.Begin()
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	// Create the new restaurant
// 	err := tx.Create(newRestaurant).Error
// 	if err != nil {
// 		tx.Rollback()
// 		if errors.Is(err, gorm.ErrDuplicatedKey) {
// 			return nil, fmt.Errorf("restaurant already exists")
// 		}
// 		return nil, err
// 	}

// 	// Create the UserRestaurant association with role 'owner'
// 	userRestaurant := &UserRestaurant{
// 		UserID:       ownerID,
// 		RestaurantID: restaurant.ID,
// 		RoleType:     "owner",
// 	}
// 	err = tx.Create(userRestaurant).Error
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	// Commit the transaction
// 	err = tx.Commit().Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return restaurant, nil
// }
