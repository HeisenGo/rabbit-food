package commands

import (
	"client/models"
	"client/services"
	"errors"
)

type LoginCommand struct {
	service services.Service
}

func (c *LoginCommand) Execute(LoginData any) (*models.User, error) {
	LoginReq, ok := LoginData.(*models.LoginUserReq)
	if !ok {
		return nil, errors.New("data type isn't LoginReq")
	}
	loggedInUser, err := c.service.Login(LoginReq)
	if err != nil {
		return nil, err
	}
	return loggedInUser, nil
}

func NewLoginCommand(service services.Service) *LoginCommand {
	return &LoginCommand{service: service}
}
