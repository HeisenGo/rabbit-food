package wallet

import (
	"context"
)

type Repo interface {
	Create(ctx context.Context, creditCard *WalletCreditCard) (*WalletCreditCard, error)
}

type WalletCreditCard struct {
	WalletID     uint
	CreditCardID uint
}

func NewWalletCreditCard(walletID uint, creditCardID uint) *WalletCreditCard {
	return &WalletCreditCard{WalletID: walletID, CreditCardID: creditCardID}
}
