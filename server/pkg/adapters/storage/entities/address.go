package entities

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Addressline  string
	Cordinates   string `gorm:"type:geography(POINT, 4326)"`
	Types        uint   //Restaurant Or User
	City         string
	UsersAddress []*User `gorm:"many2many:user_addresses;"`
	//RestaurantAddress []*User 	`gorm:"many2many:addresses_restaurants;"`
}

func NewAddressEntity() *Address {
	return &Address{}
}

type RestaurantAddress struct {
	ID           uint `gorm:"primarykey"`
	AddressID    uint
	RestaurantID uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	//*Restaurant 		*Restaurant
}
type UserAddress struct {
	ID        uint  `gorm:"primarykey"`
	AddressID uint  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignkey:AddressID;references:Address"`
	UserID    *uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      *User
}
