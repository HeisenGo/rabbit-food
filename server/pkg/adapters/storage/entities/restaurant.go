package entities

import (
	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	Name       string `gorm:"index"`
	Phone      string
	Address    *Address              `gorm:"foreignKey:RestaurantID"`
	Motors     []Motor               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // One-to-many relationship with Motor
	Categories []*RestaurantCategory `gorm:"many2many:restaurant_restaurant_categories;constraint:OnDelete:CASCADE;"`
	//Users []User `gorm:"many2many:user_restaurants;constraint:OnDelete:CASCADE;"` // Many-to-many relationship with roles
}

type UserRestaurant struct {
	gorm.Model
	UserID       uint       `gorm:"index"`
	RestaurantID uint       `gorm:"index"`
	RoleType     string     // RoleType can be 'owner', 'operator', etc.
	User         User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Restaurant   Restaurant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Menu struct {
	gorm.Model
	Name         string `gorm:"index"`
	RestaurantID uint
	Restaurant   *Restaurant `gorm:"foreignKey:RestaurantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type MenuItem struct {
	gorm.Model
	Name                          string `gorm:"index"`
	Price                         uint
	PreparationMinutes            uint
	CancellationPenaltyPercentage uint
	MenuID                        uint
	Menu                          *Menu `gorm:"foreignKey:MenuID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsAvailable                   bool  `gorm:"default:true"`
}

type RestaurantCategory struct {
	gorm.Model
	Name        string        `gorm:"index"`
	Restaurants []*Restaurant `gorm:"many2many:restaurant_restaurant_categories;constraint:OnDelete:CASCADE;"`
}

type Motor struct {
	gorm.Model
	Name         string
	Speed        int        `gorm:"index"`                                          // Speed of the motorcycle
	RestaurantID uint       `gorm:"index"`                                          // Foreign key for the restaurant
	Restaurant   Restaurant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relationship to the restaurant
}

type FunctionalAddress struct {
	City        string
	AddressLine string
}
