package services

import (
	"context"
)

type WalletService struct {
	walletOps *wallet.Ops
}

func NewWalletService(walletOps *wallet.Ops) *WalletService {
	return &WalletService{
		walletOps: walletOps,
	}
}

func (s *WalletService) CreateWallet(ctx context.Context, wallet *wallet.Wallet) (*wallet.Wallet, error) {
	createdWallet, err := s.walletOps.Create(ctx, wallet)
	if err != nil {
		return nil, err
	}
	return createdWallet, nil
}
