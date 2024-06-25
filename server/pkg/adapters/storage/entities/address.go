package entities

import (
	"gorm.io/gorm"

)

type Address struct {
	gorm.Model
	UserID      uint 			`gorm:"uniqueIndex"`
	Addressline string		
	Cordinates  string 			`gorm:"type:geography(POINT, 4326),not null"`
	Types  		uint			
	City 		string 			
	UsersAddress []*User 		`gorm:"many2many:addresses_users;"`
	RestaurantAddress []*User 	`gorm:"many2many:addresses_restaurants;"`
}
func NewAddressEntity() *Address {
	return &Address{}
}
type RestaurantAddress struct {
	ID           uint        `gorm:"primarykey"`
	UserID     	 uint        `gorm:"index:idx_together_restaurants_users,unique"`
	Address      *Address    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`        
}
type UserAddress struct {
	ID           uint        `gorm:"primarykey"`
	WalletID     uint        `gorm:"index:idx_together_wallet_card,unique"`
	Address       *Address   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`  
}