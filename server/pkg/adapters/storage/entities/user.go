package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Phone     string `gorm:"index:idx_phone,unique"`
	Email     string `gorm:"index:idx_email,unique"`
	FirstName string
	LastName  string
	Password  string
	BirthDate time.Time
	IsAdmin   bool
}
