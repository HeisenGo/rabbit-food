package entities

import "errors"

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
