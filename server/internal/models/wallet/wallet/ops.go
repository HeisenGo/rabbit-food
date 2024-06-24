package wallet

import (
	"context"
	"gorm.io/gorm"
	creditCard "server/internal/models/wallet/credit_card"
)

type WalletOps struct {
	db   *gorm.DB
	repo Repo
}

func NewWalletOps(db *gorm.DB, repo Repo) *WalletOps {
	return &WalletOps{
		db:   db,
		repo: repo,
	}
}

func (o *WalletOps) Create(ctx context.Context, wallet *Wallet) (*Wallet, error) {
	return o.repo.Create(ctx, wallet)
}

func (o *WalletOps) Deposit(ctx context.Context, creditCard *creditCard.CreditCard, amount uint) (*Wallet, error) {
	return o.repo.Deposit(ctx, creditCard, amount)
}

func (o *WalletOps) Withdraw(ctx context.Context, creditCard *creditCard.CreditCard, amount uint) (*Wallet, error) {
	return o.repo.Withdraw(ctx, creditCard, amount)
}
