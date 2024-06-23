package wallet

import (
	"context"
	"gorm.io/gorm"
)

type Ops struct {
	db   *gorm.DB
	repo Repo
}

func NewWalletOps(db *gorm.DB, repo Repo) *Ops {
	return &Ops{
		db:   db,
		repo: repo,
	}
}

func (o *Ops) Create(ctx context.Context, wallet *Wallet) (*Wallet, error) {
	return o.repo.Create(ctx, wallet)
}
