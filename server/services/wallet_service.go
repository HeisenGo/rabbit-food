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

func (s *WalletService) GetUserWalletCards(ctx context.Context) ([]*creditCard.CreditCard, error) {
	userWalletCards, err := s.creditCardOps.GetUserWalletCards(ctx)
	if err != nil {
		return nil, err
	}
	return userWalletCards, nil
}

func (s *WalletService) Deposit(ctx context.Context, card *creditCard.CreditCard, amount uint) (*wallet.Wallet, error) {
	userWallet, err := s.walletOps.Deposit(ctx, card, amount)
	if err != nil {
		return nil, err
	}
	return userWallet, nil
}

func (s *WalletService) Withdraw(ctx context.Context, card *creditCard.CreditCard, amount uint) (*wallet.Wallet, error) {
	userWallet, err := s.walletOps.Withdraw(ctx, card, amount)
	if err != nil {
		return nil, err
	}
	return userWallet, nil
}

func (s *WalletService) GetWallet(ctx context.Context, wallet *wallet.Wallet) (*wallet.Wallet, error) {
	userWallet, err := s.walletOps.GetWallet(ctx, wallet)
	if err != nil {
		return nil, err
	}
	return userWallet, nil
}
