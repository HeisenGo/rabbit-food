package models

type Wallet struct {
	ID      int
	Balance int
}

type GetWalletReq struct {
	ID int
}

type CreditCard struct {
	Number string
}
