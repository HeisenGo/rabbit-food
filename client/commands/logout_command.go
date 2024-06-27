package commands

import (
	"client/protocol/tcp"
	"client/services"
	"client/services/tcp_service"
	"errors"
)

type LogoutCommand struct {
	service services.Service
}

func (c *LogoutCommand) Execute(LogoutData any) error {
	LogoutReq, ok := LogoutData.(*tcp.LogoutUserReq)
	if !ok {
		return errors.New("data type isn't LogoutReq")
	}
	err := c.service.Logout(LogoutReq)
	if err != nil {
		return err
	}
	tcp_service.UnSetToken()
	return nil
}

func NewLogoutCommand(service services.Service) *LogoutCommand {
	return &LogoutCommand{service: service}
}
