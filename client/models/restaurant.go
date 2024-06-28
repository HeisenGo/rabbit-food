package models

type Restaurant struct {
	ID      uint     `json:"id"`
	Name    string   `json:"name"`
	Phone   string   `json:"phone"`
	Address *Address `json:"address"`
}
