package commands

import (
	"client/models"
	"client/services"
)

type DisplayCardsCommand struct {
	service services.Service
}

func (c *DisplayCardsCommand) Execute() ([]*models.CreditCard, error) {

	cards, err := c.service.DisplayCards()
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func NewDisplayCardsCommand(service services.Service) *DisplayCardsCommand {
	return &DisplayCardsCommand{service: service}
}
