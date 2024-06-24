package wallet

import (
	"context"
)

type Repo interface {
	Create(ctx context.Context, user *Wallet) (*Wallet, error)
}

type Wallet struct {
	ID      uint
	UserID  uint
	Balance uint
}

func NewWallet(userID uint) *Wallet {
	return &Wallet{
		UserID:  userID,
		Balance: 0,
	}
}
