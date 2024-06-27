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

type AddCardMenuItem struct {
	Name     string
	Command  *commands.AddCreditCardCommand
	PostMenu MenuComponent
}

func NewAddCardMenuItem(name string, command *commands.AddCreditCardCommand, postMenu MenuComponent) *AddCardMenuItem {
	return &AddCardMenuItem{
		Name:     name,
		Command:  command,
		PostMenu: postMenu,
	}
}

func (mi *AddCardMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *AddCardMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))
	var addCardData tcp.AddCardBody
	addCardData.CardNumber = utils.ReadInput(scanner, "Card Number: ")
	newCard, err := mi.Command.Execute(&addCardData)
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err, "\n")
		utils.ReadInput(scanner, "Press any key to go back... ")
		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\tCard "+newCard.Number+" Successfully added :)\n")
		utils.ReadInput(scanner, "Press any key to continue... ")
	}
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *AddCardMenuItem) GetName() string {
	return mi.Name
}
