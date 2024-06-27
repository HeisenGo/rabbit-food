package commands

import (
	"client/protocol/tcp"
	"client/services"
	"client/services/tcp_service"
	"errors"
)

type LoginCommand struct {
	service services.Service
}

func (c *LoginCommand) Execute(userData any) error {
	userReq, ok := userData.(*tcp.LoginBody)
	if !ok {
		return errors.New("data type isn't user")
	}
	token, err := c.service.Login(userReq)
	if err != nil {
		return err
	}
	tcp_service.SetToken(token.AuthorizationToken)
	return nil
}

func NewLoginCommand(service services.Service) *LoginCommand {
	return &LoginCommand{service: service}
}
