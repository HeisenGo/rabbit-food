package address

import (
	"context"
	"errors"

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

func (o *AddressOps) Delete(ctx context.Context, addressID uint) error {
	address, err := o.repo.GetByID(ctx, addressID)
	if err != nil {
		return err
	}
	if address == nil {
		return errors.New("address not found")
	}
	return o.repo.Delete(ctx, addressID)
}

func (o *AddressOps) Update(ctx context.Context, address *Address) (*Address, error) {
	existingAddress, err := o.repo.GetByID(ctx, address.UserID)
	if err != nil {
		return nil, err
	}
	if existingAddress == nil {
		return nil, errors.New("address not found")
	}
	return o.repo.Update(ctx, address)
}
