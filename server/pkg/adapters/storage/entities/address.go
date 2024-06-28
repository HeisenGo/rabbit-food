package entities

import (
	"gorm.io/gorm"
	"server/internal/models/address"
)

type Address struct {
	gorm.Model
	AddressLine string              `gorm:"size:255;not null"`
	Coordinates address.Coordinates `gorm:"type:geography(POINT, 4326);not null"`
	Types       string              `gorm:"size:255;not null"`
	City        string              `gorm:"size:255;not null"`
	UserID      uint                `gorm:"index"`
	User        User
}
