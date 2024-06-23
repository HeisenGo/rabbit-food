package entities

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	UserID      uint `gorm:"uniqueIndex"`
	Balance     uint
	CreditCards []CreditCard `gorm:"many2many:wallet_credit_cards;"`
}

type CreditCard struct {
	gorm.Model
	Number string
}

type WalletTransaction struct {
	gorm.Model
	Amount             uint
	Type               string
	Status             string
	WalletCreditCardID uint
	WalletCreditCard   *WalletCreditCard `gorm:"foreignKey:WalletCreditCardID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type WalletCreditCard struct {
	gorm.Model
	WalletID     uint `gorm:"index"`
	CreditCardID uint `gorm:"index"`
}
