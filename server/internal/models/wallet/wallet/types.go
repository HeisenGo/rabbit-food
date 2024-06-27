package wallet

import (
	"context"
	creditCard "server/internal/models/wallet/credit_card"
)

type Repo interface {
	Create(ctx context.Context, user *Wallet) (*Wallet, error)
	Deposit(ctx context.Context, creditCard *creditCard.CreditCard, amount uint) (*Wallet, error)
	Withdraw(ctx context.Context, creditCard *creditCard.CreditCard, amount uint) (*Wallet, error)
	GetWallet(ctx context.Context, wallet *Wallet) (*Wallet, error)
}

type Wallet struct {
	ID      uint
	UserID  uint
	Balance uint
}

func NewWalletByID(id uint) *Wallet {
	return &Wallet{ID: id}
}
