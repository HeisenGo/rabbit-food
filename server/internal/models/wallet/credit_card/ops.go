package wallet

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type CreditCardOps struct {
	db   *gorm.DB
	repo Repo
}

func NewCreditCardOps(db *gorm.DB, repo Repo) *CreditCardOps {
	return &CreditCardOps{
		db:   db,
		repo: repo,
	}
}

func (o *CreditCardOps) CreateCardAndAddToWallet(ctx context.Context, creditCard *CreditCard) (*CreditCard, error) {
	if !isValidCardNumber(creditCard.Number) {
		return nil, errors.New("invalid card number")
	}
	return o.repo.CreateCardAndAddToWallet(ctx, creditCard)
}

func (o *CreditCardOps) GetUserWalletCards(ctx context.Context) ([]*CreditCard, error) {
	return o.repo.GetUserWalletCards(ctx)
}
