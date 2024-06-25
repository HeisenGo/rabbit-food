package commands

import (
	"client/services"
	"fmt"
)

type DisplayCardsCommand struct {
	service services.Service
}

func (c *DisplayCardsCommand) Execute() error {

	fmt.Println("\nCards:")
	new_cards, err := c.service.DiplayCards()
	if err != nil {
		return err
	}
	for n, i := range new_cards {
		
		fmt.Println(n, i)
	}
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

func NewDisplayCardsCommand(service services.Service) *DisplayCardsCommand {
	return &DisplayCardsCommand{service: service}
}
