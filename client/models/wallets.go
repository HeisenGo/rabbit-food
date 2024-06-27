package models

type Wallet struct {
	ID      int `json:"id"`
	Balance int `json:"balance"`
}

type GetWalletReq struct {
	ID int
}

type CreditCard struct {
	ID     uint   `json:"id"`
	Number string `json:"number"`
}
