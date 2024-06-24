package commands

import (
	"client/services"
	"fmt"
)

type DisplayCreditCardsCommand struct {
	service services.Service
}

func (c *DisplayCreditCardsCommand) Execute(userData any) error {

	fmt.Println("\nBallance is:")
	// To Do
	// user, ok := userData.(*models.LoginUserReq)
	// if !ok {
	// 	return errors.New("data type isn't user")
	// }
	// token, err := c.service.Login(user)
	// fmt.Println("New: ", token)
	// fmt.Println("token: ", token.AuthorizationToken,
	// 	"\nReferesh:", token.RefreshToken,
	// 	"\nexpire: ", token.ExpiresAt)
	// //time.Sleep(time.Minute * 2)
	return nil
}

func NewDisplayCreditCardsCommand(service services.Service) *DisplayCreditCardsCommand {
	return &DisplayCreditCardsCommand{service: service}
}
