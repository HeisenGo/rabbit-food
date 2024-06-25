package storage

import (
	"context"
	"errors"
	"fmt"
	"server"
	"server/internal/errors/restaurants"
	"server/internal/models/restaurant/restaurant"
	userRestaurant "server/internal/models/restaurant/user_restaurant"
	"server/internal/models/user"
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

func (r *restaurantRepo) CreateRestaurantAndAssignOwner(ctx context.Context, restauran *restaurant.Restaurant) (*restaurant.Restaurant, error) {
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
	newRestaurantEntity := mappers.RestaurantDomainToEntity(restauran)
	err = tx.Create(newRestaurantEntity).Error
	if err != nil {
		tx.Rollback()
		// if errors.Is(err, gorm.ErrDuplicatedKey) {
		// 	return nil, fmt.Errorf("restaurant already exists")
		// }
		return nil, err
	}

	// Create the UserRestaurant association with role 'owner'
	userRestaurant := userRestaurant.NewUserRestaurant(ownerID, newRestaurantEntity.ID, server.Owner)
	err = tx.Create(userRestaurant).Error
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

	// newRestaurant := mappers.RestaurantDomainToEntity(restaurant)
	// err := r.db.Create(&newRestaurant).Error
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrDuplicatedKey) {
	// 		return nil, users.ErrUserExists
	// 	}
	// 	return nil, err
	// }
	// createdRestaurant := mappers.RestaurantEntityToDomain(newRestaurant)
	// return createdRestaurant, nil
}

func (r *restaurantRepo) GetRestaurantByID(ctx context.Context, restaurantID uint) (*restaurant.Restaurant, error) {
	var restautantEntity entities.Restaurant
	err := r.db.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ?", restaurantID).First(&restautantEntity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, restaurants.ErrRestaurantNotFound
		}
		return nil, err
	}
	return mappers.RestaurantEntityToDomain(&restautantEntity), nil
}

func (r *restaurantRepo) CheckMatchedRestaurantsOwnerIdAndClaimedID(ctx context.Context, restaurantID uint) (bool, error) {
	var userrestautantEntity entities.UserRestaurant
	err := r.db.WithContext(ctx).Model(&entities.UserRestaurant{}).Where("restaurant_id = ? AND role_type = ?", restaurantID, server.Owner).First(&userrestautantEntity).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, restaurants.ErrMismatchedOwner
		}
		return false, err
	}
	ownerID, err := utils.GetUserIDFromContext(ctx)
	if userrestautantEntity.UserID != ownerID {
		return false, restaurants.ErrMismatchedOwner
	}
	return true, nil
}

func GetAllOperators(ctx context.Context, restaurantID uint) ([]user.User, error) {
	var operators []user.User
	// err := db.WithContext(ctx).
	//     Model(&User{}).
	//     Joins("JOIN user_restaurants ON user_restaurants.user_id = users.id").
	//     Where("user_restaurants.restaurant_id = ? AND user_restaurants.role_type = ?", restaurantID, OperatorRoleType).
	//     Preload("Restaurants"). // Preload the Restaurants if needed
	//     Find(&operators).Error
	return operators, nil
}

// func (r *creditCardRepo) GetUserWalletCards(ctx context.Context) ([]*creditCard.CreditCard, error) {
// 	userID, err := utils.GetUserIDFromContext(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var creditCardEntities []*entities.CreditCard

// 	err = r.db.Joins("JOIN wallet_credit_cards ON wallet_credit_cards.credit_card_id = credit_cards.id").
// 		Joins("JOIN wallets ON wallets.id = wallet_credit_cards.wallet_id").
// 		Where("wallets.user_id = ?", userID).
// 		Find(&creditCardEntities).Error

// 	if err != nil {
// 		return nil, err
// 	}
// 	allDomainCards := mappers.BatchCreditCardEntityToDomain(creditCardEntities)
// 	return allDomainCards, nil
// }
