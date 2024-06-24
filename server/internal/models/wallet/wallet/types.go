package wallet

import (
	"context"
	creditCard "server/internal/models/wallet/credit_card"
)

type Repo interface {
	Create(ctx context.Context, user *Wallet) (*Wallet, error)
	Deposit(ctx context.Context, creditCard *creditCard.CreditCard, amount uint) (*Wallet, error)
	Withdraw(ctx context.Context, creditCard *creditCard.CreditCard, amount uint) (*Wallet, error)
}

type Wallet struct {
	ID      uint
	UserID  uint
	Balance uint
}

func NewWallet() *Wallet {
	return &Wallet{}
}
