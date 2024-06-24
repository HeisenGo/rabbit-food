package mappers

import (
	"server/internal/models/wallet/wallet"
	"server/pkg/adapters/storage/entities"
)

func WalletEntityToDomain(entity *entities.Wallet) *wallet.Wallet {
	return &wallet.Wallet{
		ID:      entity.ID,
		UserID:  entity.UserID,
		Balance: entity.Balance,
	}
}

func WalletDomainToEntity(domainWallet *wallet.Wallet) *entities.Wallet {
	return &entities.Wallet{
		UserID:  domainWallet.UserID,
		Balance: domainWallet.Balance,
	}
}
