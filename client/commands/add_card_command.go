package commands

import (
	"client/errors"
	"client/models"
	"client/protocol/tcp"
	"client/services"
)

type AddCreditCardCommand struct {
	service services.Service
}

func (c *AddCreditCardCommand) Execute(addCardData any) (*models.CreditCard, error) {
	addCardBody, ok := addCardData.(*tcp.AddCardBody)
	if !ok {
		return nil, errors.ErrDataType
	}
	addedCard, err := c.service.AddCard(addCardBody)
	if err != nil {
		return nil, err
	}
	return addedCard, nil
}

func NewAddCardCommand(service services.Service) *AddCreditCardCommand {
	return &AddCreditCardCommand{service: service}
}
