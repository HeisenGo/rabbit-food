package storage

import (
	"context"
	"errors"
	"fmt"
	"server"
	"server/internal/errors/restaurants"
	"server/internal/models/restaurant/motor"
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
		return nil, restaurants.ErrFailedRetrieveID
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

func (r *restaurantRepo) GetRestaurantByID(ctx context.Context, restaurantID uint) (*restaurant.Restaurant, error) {
	var restaurantEntity entities.Restaurant
	err := r.db.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ?", restaurantID).First(&restaurantEntity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, restaurants.ErrRestaurantNotFound
		}
		return nil, err
	}
	return mappers.RestaurantEntityToDomain(&restaurantEntity), nil
}

func (r *restaurantRepo) CheckMatchedRestaurantsOwnerIdAndClaimedID(ctx context.Context, restaurantID uint) (bool, error) {
	var userRestaurantEntity entities.UserRestaurant
	err := r.db.WithContext(ctx).Model(&entities.UserRestaurant{}).Where("restaurant_id = ? AND role_type = ?", restaurantID, server.Owner).First(&userRestaurantEntity).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, restaurants.ErrMismatchedOwner
		}
		return false, err
	}
	ownerID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return false, restaurants.ErrFailedRetrieveID
	}
	if userRestaurantEntity.UserID != ownerID {
		return false, restaurants.ErrMismatchedOwner
	}
	return true, nil
}

func (r *restaurantRepo) GetByID(ctx context.Context, restaurantID uint) (*restaurant.Restaurant, error) {
	var restaurantEntity entities.Restaurant
	err := r.db.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ?", restaurantID).First(&restaurantEntity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, restaurants.ErrRestaurantNotFound
		}
		return nil, err
	}
	return mappers.RestaurantEntityToDomain(&restaurantEntity), nil
}

func (r *restaurantRepo) AssignOperatorToRestaurant(ctx context.Context, operator *user.User, restaurant restaurant.Restaurant) (*user.User, error) {
	userRestaurant := userRestaurant.NewUserRestaurant(operator.ID, restaurant.ID, server.Operator)
	err := r.db.Create(&userRestaurant).Error
	if err != nil {
		return nil, restaurants.ErrOperatorAssignFailed
	}
	return operator, nil
}

func (r *restaurantRepo) RemoveOperatorFromRestaurant(ctx context.Context, operatorID uint, restaurantID uint) error {
	err := r.db.Where("role_type = ? AND restaurant_id=? AND user_id=?", server.Operator, restaurantID, operatorID).Delete(&entities.UserRestaurant{}).Error
	if err != nil {
		return restaurants.ErrRemoveOperatorFailed
	}
	return nil
}

func (r *restaurantRepo) WithdrawRestaurant(ctx context.Context, newOwnerID uint, restaurantID uint) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
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
		return restaurants.ErrFailedRetrieveID
	}
	// Create the new restaurant
	err = tx.Where("user_id = ? AND restaurant_id = ? AND role_type=?", ownerID, restaurantID, server.Owner).Delete(&entities.UserRestaurant{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Create the new UserRestaurant association with role 'owner'
	newUserRestaurantEntity := userRestaurant.NewUserRestaurant(newOwnerID, restaurantID, server.Owner)
	err = tx.Create(&newUserRestaurantEntity).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	err = tx.Commit().Error
	if err != nil {
		return err
	}
	return nil
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

func (r *restaurantRepo) AddMotor(ctx context.Context, motor *motor.Motor, restaurantID uint) (*motor.Motor, error) {
	// userRestaurant := userRestaurant.NewUserRestaurant(operator.ID, restaurant.ID, server.Operator)
	motorEntity := mappers.MotorDomainToEntity(motor)
	motorEntity.RestaurantID = restaurantID
	err := r.db.Create(&motorEntity).Error
	if err != nil {
		return nil, restaurants.ErrMotorAdditionFailed
	}
	return mappers.MotorEntityToDomain(motorEntity), nil
}

func (r *restaurantRepo) RemoveMotor(ctx context.Context, motorID uint) error {
	err := r.db.Delete(&entities.Motor{}, motorID).Error
	if err != nil {
		return restaurants.ErrRemoveOperatorFailed
	}
	return nil
}

func (r *restaurantRepo) GetAllMotors(ctx context.Context, restaurantID uint) ([]*motor.Motor, error) {
	// userID, err := utils.GetUserIDFromContext(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	var restaurant entities.Restaurant
	err := r.db.Preload("Motors").First(&restaurant, restaurantID).Error
	if err != nil {
		return nil, err
	}
	motors := []*motor.Motor{}
	for _, motor := range restaurant.Motors {
		domainMotor := mappers.MotorEntityToDomain(&motor)
		motors = append(motors, domainMotor)
	}
	return motors, nil
}

func (r *restaurantRepo) GetAllOperators(ctx context.Context, restaurantID uint) ([]*user.User, error) {
	var operators []*entities.User

	err := r.db.Joins("JOIN user_restaurants ON user_restaurants.user_id = users.id").
		Where("user_restaurants.restaurant_id = ? AND user_restaurants.role_type = ?", restaurantID, server.Operator).
		Find(&operators).Error

	if err != nil {
		return nil, err
	}

	domainOperators := []*user.User{}
	for _, user := range operators {
		dUser := mappers.UserEntityToDomain(user)
		domainOperators = append(domainOperators, dUser)
	}
	return domainOperators, nil
}

func (r *restaurantRepo) GetRestaurantsToAddCategoryMenuFood(ctx context.Context) ([]*restaurant.Restaurant, error) {
	var restaurants []entities.Restaurant
	workingUserID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.New("failed to recognize the working user")
	}
	err = r.db.WithContext(ctx).
		Joins("JOIN user_restaurants ON user_restaurants.restaurant_id = restaurants.id").
		Where("user_restaurants.user_id = ? AND user_restaurants.role_type IN ?", workingUserID, []string{string(server.Owner), string(server.Operator)}).
		Distinct("restaurants.id, restaurants.name, restaurants.phone").
		Find(&restaurants).Error

	if err != nil {
		return nil, errors.New("internal error")
	}
	domainRestaurants := []*restaurant.Restaurant{}
	for _, rest := range restaurants {
		//var address entities.FunctionalAddress
		sqlStr := "SELECT city, address_line FROM addresses WHERE restaurant_id = ?"
		err = r.db.WithContext(ctx).Raw(sqlStr, rest.ID).Scan(&rest.Address).Error
		if err != nil {
			fmt.Println("Error retrieving restaurant:", err)
		}
		dRest := mappers.RestaurantEntityAddressNameLineToDomain(&rest)
		domainRestaurants = append(domainRestaurants, dRest)
	}
	return domainRestaurants, nil

}

func (r *restaurantRepo) DoesThisHaveARoleInRestaurant(ctx context.Context, restaurantID uint) (bool, error) {
	var restaurantUsers []*entities.UserRestaurant

	err := r.db.WithContext(ctx).
		Joins("JOIN user_restaurants ON user_restaurants.user_id = users.id").
		Where("user_restaurants.restaurant_id = ? AND user_restaurants.role_type IN ?", restaurantID, []string{string(server.Operator), string(server.Owner)}).
		Find(&restaurantUsers).Error
	if err != nil {
		return false, err
	}

	workingUserID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return false, err
	}

	for _, user := range restaurantUsers {
		if user.ID == workingUserID {
			return true, nil
		}
	}

	return false, restaurants.ErrUserNotAllowed
}

func (r *restaurantRepo) GetOwnerInfo(ctx context.Context, restaurantID uint) (*user.User, error) {
	var userRestaurant entities.UserRestaurant
	err := r.db.Preload("User").Where("restaurant_id = ? AND role_type = ?", restaurantID, "owner").First(&userRestaurant).Error
	if err != nil {
		return nil, err
	}

	owner := userRestaurant.User
	return mappers.UserEntityToDomain(&owner), nil
}

func (r *restaurantRepo) GetRestaurantInfo(ctx context.Context, restaurantID uint) (*restaurant.Restaurant, *user.User, []*user.User, []*motor.Motor, error) {
	restaurant, err := r.GetByID(ctx, restaurantID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	owner, err := r.GetOwnerInfo(ctx, restaurantID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	operators, err := r.GetAllOperators(ctx, restaurantID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	motors, err := r.GetAllMotors(ctx, restaurantID)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return restaurant, owner, operators, motors, nil
}

func (r *restaurantRepo) EditRestaurantName(ctx context.Context, restaurantID uint, newName string) error {
	err := r.db.Model(&entities.Restaurant{}).Where("id = ?", restaurantID).Update("Name", newName).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantRepo) RemoveRestaurant(ctx context.Context, restaurantID uint) error {
	//err := r.db.Where("role_type = ? AND restaurant_id=? AND user_id=?", server.Operator, restaurantID, operatorID).Delete(&entities.UserRestaurant{}).Error
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if rv := recover(); rv != nil {
			tx.Rollback()
		}
	}()
	// delete the association of restaurant
	err := tx.Where("restaurant_id = ?", restaurantID).Delete(&entities.UserRestaurant{}).Error
	if err != nil {
		tx.Rollback()
		// if errors.Is(err, gorm.ErrDuplicatedKey) {
		// 	return nil, fmt.Errorf("restaurant already exists")
		// }
		return err
	}

	//** delete restaurant foods/menu/category

	// delete the Restaurant
	err = tx.Where("restaurant_id = ?", restaurantID).Delete(&entities.Restaurant{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	err = tx.Commit().Error
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantRepo) GetRestaurantsOfAnOwner(ctx context.Context) ([]*restaurant.Restaurant, error) {

	var owningRestaurants []*entities.Restaurant
	ownerID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		//logger?
		fmt.Println("UserId could not be recognized in context to create a restaurant for it")
		return nil, restaurants.ErrFailedRetrieveID
	}

	err = r.db.Joins("JOIN user_restaurants ON user_restaurants.restaurant_id = restaurants.id").
		Where("user_restaurants.user_id = ? AND user_restaurants.role_type = ?", ownerID, server.Owner).
		Find(&owningRestaurants).Error

	// err = r.db.Joins("JOIN user_restaurants ON user_restaurants.restaurant_id = restaurant.id").
	// 	Joins("JOIN restaurants ON restaurant.id = ").Where("user_restaurants.restaurant_id = ? AND user_restaurants.role_type = ?", ownerID, server.Owner).
	// 	Find(&restaurants).Error

	if err != nil {
		return nil, err
	}

	domainRestaurants := []*restaurant.Restaurant{}
	for _, rest := range owningRestaurants {
		//var address entities.FunctionalAddress
		sqlStr := "SELECT city, address_line FROM addresses WHERE restaurant_id = ?"
		err = r.db.WithContext(ctx).Raw(sqlStr, rest.ID).Scan(&rest.Address).Error
		if err != nil {
			fmt.Println("Error retrieving restaurant:", err)
		}
		dRest := mappers.RestaurantEntityAddressNameLineToDomain(rest)
		domainRestaurants = append(domainRestaurants, dRest)
	}
	return domainRestaurants, nil

}

func (r *restaurantRepo) GetRestaurantsOfAnOperator(ctx context.Context) ([]*restaurant.Restaurant, error) {

	var operatingRestaurants []*entities.Restaurant
	operatorID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		//logger?
		fmt.Println("UserId could not be recognized in context to create a restaurant for it")
		return nil, restaurants.ErrFailedRetrieveID
	}

	err = r.db.Joins("JOIN user_restaurants ON user_restaurants.restaurant_id = restaurants.id").
		Where("user_restaurants.user_id = ? AND user_restaurants.role_type = ?", operatorID, server.Operator).
		Find(&operatingRestaurants).Error

	// err = r.db.Joins("JOIN user_restaurants ON user_restaurants.restaurant_id = restaurant.id").
	// 	Joins("JOIN restaurants ON restaurant.id = ").Where("user_restaurants.restaurant_id = ? AND user_restaurants.role_type = ?", ownerID, server.Owner).
	// 	Find(&restaurants).Error

	if err != nil {
		return nil, err
	}

	domainOperatingRestaurants := []*restaurant.Restaurant{}
	for _, rest := range operatingRestaurants {
		dRest := mappers.RestaurantEntityToDomain(rest)
		domainOperatingRestaurants = append(domainOperatingRestaurants, dRest)
	}
	return domainOperatingRestaurants, nil

}

/// func (r *restaurantRepo) EditRestaurantPhone() error

/// func (r *restaurantRepo) EditRestaurantAddress() error

/// func (r *restaurantRepo) GetRestaurantAddress()

/// func (r *restaurantRepo) GetRestaurantCategories()

///// func (r *restaurantRepo) GetRestaurantMenus()

///// func (r *restaurantRepo) GetRestaurantFoods()

// func (r *restaurantRepo) RemoveRestaurant()
