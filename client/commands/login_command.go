package commands

import (
	"client/protocol/tcp"
	"client/services"
	"client/services/tcp_service"
	"errors"
	"fmt"
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
	tcp_service.SetToken(token.AuthorizationToken)
	fmt.Println("New: ", token)
	fmt.Println("token: ", token.AuthorizationToken,
		"\nReferesh:", token.RefreshToken,
		"\nexpire: ", token.ExpiresAt)
	//time.Sleep(time.Minute * 2)
	return err
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

func NewLoginCommand(service services.Service) *LoginCommand {
	return &LoginCommand{service: service}
}
