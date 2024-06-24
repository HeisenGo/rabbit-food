package commands

import (
	"client/models"
	"client/services"
	"errors"
)

type LogoutCommand struct {
	service services.Service
}

func (c *LogoutCommand) Execute(LogoutData any) error {
	LogoutReq, ok := LogoutData.(*models.LogoutUserReq)
	if !ok {
		return errors.New("data type isn't LogoutReq")
	}
	err := c.service.Logout(LogoutReq)
	if err != nil {
		return err
	}
	return nil
}

func NewLogoutCommand(service services.Service) *LogoutCommand {
	return &LogoutCommand{service: service}
}
