package storage

import (
	"context"
	"gorm.io/gorm"
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
		return nil, err
	}
	createdWallet := mappers.WalletEntityToDomain(newWallet)
	return createdWallet, nil
}
