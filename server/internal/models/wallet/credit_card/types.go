package wallet

import (
	"context"
)

type Repo interface {
	CreateCardAndAddToWallet(ctx context.Context, creditCard *CreditCard) (*CreditCard, error)
}

type CreditCard struct {
	ID     uint
	Number string
}

func NewCreditCard(number string) *CreditCard {
	return &CreditCard{
		Number: number,
	}
}
