package storage

import (
	"fmt"
	"server/config"
	"server/pkg/adapters/storage/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresGormConnection(dbConfig config.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		dbConfig.Host, dbConfig.User, dbConfig.Pass, dbConfig.DBName, dbConfig.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func Migrate(db *gorm.DB) error {
	migrator := db.Migrator()

	err := migrator.AutoMigrate(&entities.User{},
		&entities.Wallet{}, &entities.CreditCard{},
		&entities.WalletTransaction{},
		&entities.WalletCreditCard{},
		&entities.Restaurant{},
		&entities.UserRestaurant{},
		&entities.Menu{},
		&entities.MenuItem{},
		&entities.RestaurantCategory{},
    	&entities.Address{},
    	&entities.RestaurantAddress{},
    	&entities.UserAddress{},
	)
	if err != nil {
		return err
	}
	return nil
}
