package storage

import (
	"context"
	"fmt"
	"server"
	"server/internal/models/restaurant/restaurant"
	userRestaurant "server/internal/models/restaurant/user_restaurant"
	"server/pkg/adapters/storage/entities"
	"server/pkg/adapters/storage/mappers"
	"server/pkg/utils"

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

func (r *restaurantRepo) CreateRestaurantAndAssignOwner(ctx context.Context, restaurant *restaurant.Restaurant) (*restaurant.Restaurant, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if rv := recover(); rv != nil {
			tx.Rollback()
		}
	}()
	ownerID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		//logger?
		fmt.Println("UserId could not be recognized in context to create a restaurant for it")
		return nil, err
	}
	// Create the new restaurant
	newRestaurantEntity := mappers.RestaurantDomainToEntity(restaurant)
	err = tx.Create(newRestaurantEntity).Error
	if err != nil {
		tx.Rollback()
		// if errors.Is(err, gorm.ErrDuplicatedKey) {
		// 	return nil, fmt.Errorf("restaurant already exists")
		// }
		return nil, err
	}

	// Create the UserRestaurant association with role 'owner'
	userRest := userRestaurant.NewUserRestaurant(ownerID, newRestaurantEntity.ID, server.Owner)
	err = tx.Create(userRest).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}
	createdRestaurant := mappers.RestaurantEntityToDomain(newRestaurantEntity)

	return createdRestaurant, nil
}

func (r *restaurantRepo) AddCategoriesToRestaurant(ctx context.Context, rest *restaurant.Restaurant, categoryIDs []uint) (*restaurant.Restaurant, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if rv := recover(); rv != nil {
			tx.Rollback()
		}
	}()

	// Validate restaurant exists
	var restaurantEntity *entities.Restaurant
	err := tx.First(&restaurantEntity, rest.ID).Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("restaurant not found: %w", err)
	}

	for _, categoryID := range categoryIDs {
		restaurantCategory := entities.RestaurantCategory{
			Model: gorm.Model{ID: categoryID},
		}
		err = tx.Model(&restaurantEntity).Association("Categories").Append(&restaurantCategory)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to associate category ID %d with restaurant: %w", categoryID, err)
		}
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	updatedRestaurant := mappers.RestaurantEntityToDomain(restaurantEntity)
	return updatedRestaurant, nil
}

func (r *restaurantRepo) GetRestaurantCategories(ctx context.Context, restaurantID uint) ([]*restaurant.RestaurantCategory, error) {
	var categoryEntities []*restaurant.RestaurantCategory
	if restaurantID == 0 {
		err := r.db.Find(&categoryEntities).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := r.db.Joins("JOIN restaurant_restaurant_categories ON restaurant_restaurant_categories.restaurant_category_id = restaurant_categories.id").
			Where("restaurant_restaurant_categories.restaurant_id = ?", restaurantID).
			Find(&categoryEntities).Error
		if err != nil {
			return nil, err
		}
	}
	return categoryEntities, nil
}
