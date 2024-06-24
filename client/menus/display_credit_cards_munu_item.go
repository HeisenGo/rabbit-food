package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/utils"
	"fmt"
	"time"
)

type DisplayCreditCardsMenuItem struct {
	Name     string
	Command  *commands.DisplayCreditCardsCommand
	PostMenu MenuComponent
}

func NewDisplayCreditCardsMenuItem(name string, command *commands.DisplayCreditCardsCommand, postMenu MenuComponent) *DisplayCreditCardsMenuItem {
	return &DisplayCreditCardsMenuItem{
		Name:     name,
		Command:  command,
		PostMenu: postMenu,
	}
}

func (mi *DisplayCreditCardsMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *DisplayCreditCardsMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))

	///////////////////// To do

	/////////////////////////
	// TODO: mi.state = ProfileMngmnt state
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *DisplayCreditCardsMenuItem) GetName() string {
	return mi.Name
}
