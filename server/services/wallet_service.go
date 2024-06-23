package services

import (
	"context"
	"server/internal/models/wallet/wallet"
)

type WalletService struct {
	walletOps *wallet.Ops
}

func NewWalletService(walletOps *wallet.Ops) *WalletService {
	return &WalletService{
		walletOps: walletOps,
	}
}

func (s *WalletService) CreateWalletByUserID(ctx context.Context, userID uint) (*wallet.Wallet, error) {
	newWallet := wallet.NewWallet(userID)
	createdWallet, err := s.walletOps.Create(ctx, newWallet)
	if err != nil {
		return nil, err
	}
	return createdWallet, nil
}
