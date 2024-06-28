package entities

import (
	"errors"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	RestaurantID   uint
	Restaurant     *Restaurant `gorm:"foreignkey:RestaurantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User           *User       `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID         uint
	OrderStatus    OrderStatus     `gorm:"type:enum('not_started','in_progress','in_delivery', 'received');not null;default:'not_started'"`
	PaymentStatus  PaymentStatus   `gorm:"type:enum('paid','unpaid');not null;default:'unpaid'"`
	FollowUpStatus *FollowUpStatus `gorm:"type:enum('in_progress','completed');"`
	TotalPrice     string
	MenuItem       []*MenuItem `gorm:"many2many:order_items;"`
}

type OrderItem struct {
	gorm.Model
	MenuItemID uint      `gorm:"index:idx_together_order_item,unique"`
	MenuItem   *MenuItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrderID    uint      `gorm:"index:idx_together_order_item,unique"`
	Order      *Order    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price      *uint
}

type OrderStatus string

const (
	OrderStatusNotStarted OrderStatus = "not_started"
	OrderStatusInProgress OrderStatus = "in_progress"
	OrderStatusInDelivery OrderStatus = "in_delivery"
	OrderStatusReceived   OrderStatus = "received"
	OrderStatusCanceled   OrderStatus = "canceled"
)

func (status OrderStatus) IsValid() error {
	switch status {
	case OrderStatusNotStarted, OrderStatusInProgress, OrderStatusInDelivery, OrderStatusReceived, OrderStatusCanceled:
		return nil
	}
	return errors.New("invalid order status")
}

type PaymentStatus string

const (
	PaymentStatusUnpaid PaymentStatus = "unpaid"
	PaymentStatusPaid   PaymentStatus = "paid"
)

func (status PaymentStatus) IsValid() error {
	switch status {
	case PaymentStatusUnpaid, PaymentStatusPaid:
		return nil
	}
	return errors.New("invalid payment status")
}

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

type FollowUpStatus string

const (
	FollowUpStatusInProgress FollowUpStatus = "unpaid"
	FollowUpStatusCompleted  FollowUpStatus = "paid"
)

func (status FollowUpStatus) IsValid() error {
	switch status {
	case FollowUpStatusInProgress, FollowUpStatusCompleted:
		return nil
	}
	return errors.New("invalid follow up status")
}
