package commands

import (
	"client/services"
)

type ConnectACreditCardCommand struct {
	service services.Service
}

func (c *ConnectACreditCardCommand) Execute(userData any) error {

	//time.Sleep(time.Minute * 2)
	return nil
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

func NewConnectACreditCardCommand(service services.Service) *ConnectACreditCardCommand {
	return &ConnectACreditCardCommand{service: service}
}
