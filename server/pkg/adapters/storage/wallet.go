package storage

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"server/internal/errors/users"
	"server/internal/models/wallet/wallet"
	"server/pkg/adapters/storage/mappers"
)

type walletRepo struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) wallet.Repo {
	return &walletRepo{
		db: db,
	}
}
func (r *walletRepo) Create(ctx context.Context, wallet *wallet.Wallet) (*wallet.Wallet, error) {
	newWallet := mappers.WalletDomainToEntity(wallet)
	err := r.db.Create(&newWallet).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, users.ErrUserExists
		}
		return nil, err
	}
	createdWallet := mappers.WalletEntityToDomain(newWallet)
	return createdWallet, nil
}
