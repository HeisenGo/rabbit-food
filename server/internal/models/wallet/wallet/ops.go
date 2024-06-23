package wallet

import (
	"context"
	"gorm.io/gorm"
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
