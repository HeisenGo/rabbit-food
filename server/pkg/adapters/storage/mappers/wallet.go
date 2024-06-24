package mappers

import (
	creditCard "server/internal/models/wallet/credit_card"
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

func CreditCardEntityToDomain(entity *entities.CreditCard) *creditCard.CreditCard {
	return &creditCard.CreditCard{
		ID:     entity.ID,
		Number: entity.Number,
	}
}

func BatchCreditCardEntityToDomain(entities []*entities.CreditCard) []*creditCard.CreditCard {
	var domainCreditCards []*creditCard.CreditCard
	for _, e := range entities {
		domainCreditCards = append(domainCreditCards, &creditCard.CreditCard{ID: e.ID, Number: e.Number})
	}
	return domainCreditCards
}

func CreditCardDomainToEntity(domainWallet *creditCard.CreditCard) *entities.CreditCard {
	return &entities.CreditCard{
		Number: domainWallet.Number,
	}
}
