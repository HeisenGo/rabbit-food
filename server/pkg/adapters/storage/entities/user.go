package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Phone     string
	Email     string
	FirstName string
	LastName  string
	Password  string
	BirthDate time.Time
	IsAdmin		bool
}
