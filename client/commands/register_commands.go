package commands

import (
	"client/errors"
	"client/models"
	"client/services"
	"client/services/tcp_service"
)

type RegisterCommand struct {
	service services.Service
}

func (c *RegisterCommand) Execute(userData any) error {
	user, ok := userData.(*models.User)
	if !ok {
		return errors.ErrDataType
	}
	token, err := c.service.Register(user)
	if err != nil {
		return err
	}
	tcp_service.SetToken(token.AuthorizationToken)
	return nil
}

func NewRegisterCommand(service services.Service) *RegisterCommand {
	return &RegisterCommand{service: service}
}
