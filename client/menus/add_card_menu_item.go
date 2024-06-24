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

// func (mi *LoginMenuItem) Execute(scanner *bufio.Scanner) {
// 	defer time.Sleep(time.Second)
// 	utils.ClearScreen()
// 	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))
// }

func (mi *AddCardMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))
	var addCardData tcp.AddCardBody
	addCardData.CardNumber = utils.ReadInput(scanner, "Card Number: ")
	_, err := mi.Command.Execute(&addCardData)
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

func (mi *AddCardMenuItem) GetName() string {
	return mi.Name
}
