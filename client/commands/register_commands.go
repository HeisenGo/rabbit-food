package commands

import (
	"client/models"
	"client/services"
	"client/services/tcp_service"
	"errors"
	"fmt"
)

type RegisterCommand struct {
	service services.Service
}

func (c *RegisterCommand) Execute(userData any) error {
	user, ok := userData.(*models.User)
	if !ok {
		return errors.New("data type isn't user")
	}
	token, err := c.service.Register(user)
	tcp_service.SetToken(token.AuthorizationToken)
	if err != nil {
		return err
	}
	fmt.Println("New: ", token)
	fmt.Println("token: ", token.AuthorizationToken,
		"\nReferesh:", token.RefreshToken,
		"\nexpire: ", token.ExpiresAt)
	//time.Sleep(time.Minute * 2)

	return err
}

func NewRegisterCommand(service services.Service) *RegisterCommand {
	return &RegisterCommand{service: service}
}
