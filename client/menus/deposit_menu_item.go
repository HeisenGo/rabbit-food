package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/protocol/tcp"
	"client/utils"
	"fmt"
	"time"
)

type DepositCardMenuItem struct {
	Name                  string
	DepositCommand        *commands.DepositCommand
	GetWalletCardsCommand *commands.GetWalletCardsCommand
	PostMenu              MenuComponent
}

func NewDepositMenuItem(name string, depositCommand *commands.DepositCommand, getWalletCardsCommand *commands.GetWalletCardsCommand, postMenu MenuComponent) *DepositCardMenuItem {
	return &DepositCardMenuItem{
		Name:                  name,
		DepositCommand:        depositCommand,
		GetWalletCardsCommand: getWalletCardsCommand,
		PostMenu:              postMenu,
	}
}

func (mi *DepositCardMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *DepositCardMenuItem) Execute(scanner *bufio.Scanner) {
	cards, err := mi.GetWalletCardsCommand.Execute()
	for i, c := range cards {
		utils.ColoredPrint(constants.Purple, fmt.Sprintf("%d- %v", i, c.Number))
	}
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))
	var addCardData tcp.AddCardBody
	addCardData.CardNumber = utils.ReadInput(scanner, "Card Number: ")
	_, err := mi.DepositCommand.Execute(&addCardData)
	if err != nil {
		fmt.Println(err)
		utils.ReadInput(scanner, "Press any key to continue... ")

		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\tSuccessful Login :)\n")
		utils.ReadInput(scanner, "Press any key to continue... ")
	}
	// TODO: mi.state = ProfileMngmnt state
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *AddCreditCardMenuItem) GetName() string {
	return mi.Name
}
