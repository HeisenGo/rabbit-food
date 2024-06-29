package entities

import (
	"gorm.io/gorm"
)

func (order *Order) BeforeSave(tx *gorm.DB) (err error) {
	if err := order.PaymentStatus.IsValid(); err != nil {
		return err
	}
	if err := order.OrderStatus.IsValid(); err != nil {
		return err
	}
	if err := order.FollowUpStatus.IsValid(); err != nil {
		return err
	}
	return nil
}
