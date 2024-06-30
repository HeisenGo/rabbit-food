package models

type Restaurant struct {
	ID      uint     `json:"id"`
	Name    string   `json:"name"`
	Phone   string   `json:"phone"`
	Address *Address `json:"address"`
}

type MenuItem struct {
	ID                            uint   `json:"id"`
	Name                          string `json:"name"`
	Price                         uint   `json:"price"`
	PreparationMinutes            uint   `json:"preparation_minutes"` // in minutes
	CancellationPenaltyPercentage uint   `json:"cancellation_penalty_percentage"`
	MenuID                        uint   `json:"menu_id"`
}
