package commands

import (
	"client/models"
	"client/services"
	"errors"
)

type RegisterCommand struct {
	service services.Service
}

func (c *RegisterCommand) Execute(userData any) error {
	user, ok := userData.(*models.User)
	if !ok {
		return errors.New("data type isn't user")
	}
	err := c.service.Register(user)
	return err
}

func NewRegisterCommand(service services.Service) *RegisterCommand {
	return &RegisterCommand{service: service}
}
