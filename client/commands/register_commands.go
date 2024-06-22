package commands

import (
	"client/models"
	"client/services"
	"errors"
	"fmt"
	"time"
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
	fmt.Println("New: ", token)
	fmt.Println("token: ", token.AuthorizationToken,
		"\nReferesh:", token.RefreshToken,
		"\nexpire: ", token.ExpiresAt)
	time.Sleep(time.Minute * 2)
	return err
}

func NewRegisterCommand(service services.Service) *RegisterCommand {
	return &RegisterCommand{service: service}
}
