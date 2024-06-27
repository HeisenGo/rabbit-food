package entities

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Addressline  string `gorm:"size:255;not null"`
	Cordinates   [2]float64 `gorm:"type:geography(POINT, 4326);not null"`
	Types        string `gorm:"size:255;not null"`
	City         string `gorm:"size:255;not null"`
	UserID       *uint  `gorm:"unique;index"`
	User         User
	RestaurantID *uint `gorm:"unique;index"`
	//Restaurant   Restaurant
}

func NewAddressEntity() *Address {
	return &Address{}
}

type RestaurantAddress struct {
	ID           uint `gorm:"primarykey"`
	AddressID    uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignkey:AddressID;references:address"`
	RestaurantID uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignkey:RestaurantID;references:restaurant"`
	//*Restaurant 		*Restaurant
}
type UserAddress struct {
	ID        uint  `gorm:"primarykey"`
	AddressID uint  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignkey:AddressID;references:address"`
	UserID    *uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignkey:UserID;references:users"`
	User      *User
}
