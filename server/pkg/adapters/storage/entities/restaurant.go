package entities

import (
	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:user_restaurants;constraint:OnDelete:CASCADE;"` // Many-to-many relationship with roles
}

type UserRestaurant struct {
	gorm.Model
	UserID       uint       `gorm:"index"`
	RestaurantID uint       `gorm:"index"`
	RoleType     string     // RoleType can be 'owner', 'operator', etc.
	User         User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Restaurant   Restaurant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
