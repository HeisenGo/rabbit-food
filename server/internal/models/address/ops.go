package address

import (
	"context"
	"gorm.io/gorm"
)

type AddressOps struct {
	db   *gorm.DB
	repo Repo
}

func NewAddressOps(db *gorm.DB, repo Repo) *AddressOps {
	return &AddressOps{
		db:   db,
		repo: repo,
	}
}

func (o *AddressOps) Create(ctx context.Context, address *Address) (*Address, error) {
	return o.repo.Create(ctx, address)
}

