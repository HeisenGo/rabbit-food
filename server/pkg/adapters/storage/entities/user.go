package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Phone     string  `gorm:"index:idx_phone,unique"`
	Email     *string `gorm:"uniqueIndex:idx_email_not_null,where:email IS NOT NULL"`
	FirstName string
	LastName  string
	Password  string
	BirthDate time.Time
	IsAdmin   bool    `gorm:"default:false"`
	Wallet    *Wallet `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
