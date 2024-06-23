package mappers

import (
	wallet2 "server/internal/models/wallet/credit_card"
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

func CreditCardEntityToDomain(entity *entities.CreditCard) *wallet2.CreditCard {
	return &wallet2.CreditCard{
		ID:     entity.ID,
		Number: entity.Number,
	}
}

func CreditCardDomainToEntity(domainWallet *wallet2.CreditCard) *entities.CreditCard {
	return &entities.CreditCard{
		Number: domainWallet.Number,
	}
}
