package entities

import (
	"gorm.io/gorm"
)

type Address struct {
    gorm.Model
    Addressline  string     `gorm:"size:255;not null"`
    Cordinates  [2]float64 `gorm:"type:geography(POINT, 4326);not null"`
    Types        string     `gorm:"size:255;not null"`
    City         string     `gorm:"size:255;not null"`
    UserID       *uint      `gorm:"index"` // remove unique constraint for one-to-many relationship
    User         User
    RestaurantID *uint      `gorm:"index"` // remove unique constraint for one-to-many relationship
    //Restaurant   Restaurant
}

func NewAddressEntity() *Address {
	return &Address{}
}

type RestaurantAddress struct {
    ID           uint `gorm:"primarykey"`
    AddressID    uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AddressID;references:ID"`
    RestaurantID uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RestaurantID;references:ID"`
    //Restaurant   Restaurant
}
type UserAddress struct {
    ID        uint  `gorm:"primarykey"`
    AddressID uint  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AddressID;references:ID"`
    UserID    *uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID;references:ID"`
    User      User
}