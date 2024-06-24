package commands

import (
	"client/models"
	"client/protocol/tcp"
	"client/services"
	"errors"
)

type AddCreditCardCommand struct {
	service services.Service
}

func (c *AddCreditCardCommand) Execute(addCardData any) (*models.CreditCard, error) {
	addCardBody, ok := addCardData.(*tcp.AddCardBody)
	if !ok {
		return nil, errors.New("data type isn't AddCardReq")
	}
	addedCard, err := c.service.AddCard(addCardBody)
	if err != nil {
		return nil, err
	}
	return addedCard, nil
}

// func (c *LoginCommand) Execute(LoginData any) (*models.User, error) {
// 	LoginReq, ok := LoginData.(*models.LoginUserReq)
// 	if !ok {
// 		return nil, errors.New("data type isn't LoginReq")
// 	}
// 	loggedInUser, err := c.service.Login(LoginReq)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return loggedInUser, nil
// }

func NewAddCardCommand(service services.Service) *AddCreditCardCommand {
	return &AddCreditCardCommand{service: service}
}
