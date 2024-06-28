package entities

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	RestaurantID            uint
	Restaurant              *Restaurant `gorm:"foreignkey:RestaurantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User                    *User       `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID                  uint
	OrderStatus             OrderStatus   `gorm:"default:'not_started'"`
	PaymentStatus           PaymentStatus `gorm:"default:'unpaid'"`
	FollowUpStatus          *FollowUpStatus
	TotalPrice              string
	TotalPreparationMinutes uint
	PreparationStartAt      time.Time
	MenuItem                []*MenuItem `gorm:"many2many:order_items;"`
}

type OrderItem struct {
	gorm.Model
	MenuItemID uint      `gorm:"index:idx_together_order_item,unique"`
	MenuItem   *MenuItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrderID    uint      `gorm:"index:idx_together_order_item,unique"`
	Order      *Order    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price      *uint
}
