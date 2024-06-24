package services

import (
	"context"
	creditCard "server/internal/models/wallet/credit_card"
	"server/internal/models/wallet/wallet"
)

type WalletService struct {
	walletOps     *wallet.WalletOps
	creditCardOps *creditCard.CreditCardOps
}

func NewWalletService(walletOps *wallet.WalletOps, creditCardOps *creditCard.CreditCardOps) *WalletService {
	return &WalletService{
		walletOps:     walletOps,
		creditCardOps: creditCardOps,
	}
}

func (s *WalletService) AddCardToWalletByUserID(ctx context.Context, card *creditCard.CreditCard) (*creditCard.CreditCard, error) {
	createdCard, err := s.creditCardOps.CreateCardAndAddToWallet(ctx, card)
	if err != nil {
		return nil, err
	}
	return createdCard, nil
}
