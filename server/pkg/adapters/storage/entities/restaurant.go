package entities

import (
	"gorm.io/gorm"
)



type Restaurant struct {
	gorm.Model
	Name   string `gorm:"index"`
	Phone  string
	Motors []Motor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // One-to-many relationship with Motor
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

type Motor struct {
	gorm.Model
	Name         string
	Speed        int        `gorm:"index"`                                          // Speed of the motorcycle
	RestaurantID uint       `gorm:"index"`                                          // Foreign key for the restaurant
	Restaurant   Restaurant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Relationship to the restaurant
}
