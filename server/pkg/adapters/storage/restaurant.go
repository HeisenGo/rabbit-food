package storage

import (
	"context"
	"fmt"
	"server"
	"server/internal/models/restaurant/restaurant"
	userRestaurant "server/internal/models/restaurant/user_restaurant"
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
