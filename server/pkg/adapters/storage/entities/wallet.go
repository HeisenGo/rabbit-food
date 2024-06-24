package entities

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	UserID      uint `gorm:"uniqueIndex"`
	Balance     uint
	CreditCards []*CreditCard `gorm:"many2many:wallet_credit_cards;"`
}

func NewWalletEntity() *Wallet {
	return &Wallet{}
}

type CreditCard struct {
	gorm.Model
	Number  string    `gorm:"uniqueIndex"`
	Wallets []*Wallet `gorm:"many2many:wallet_credit_cards;"`
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
	ID           uint        `gorm:"primarykey"`
	WalletID     uint        `gorm:"index:idx_together_wallet_card,unique"`
	Wallet       *Wallet     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreditCardID uint        `gorm:"index:idx_together_wallet_card,unique"`
	CreditCard   *CreditCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
