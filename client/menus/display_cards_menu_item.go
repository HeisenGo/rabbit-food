package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/utils"
	"fmt"
	"time"
)

type DisplayCardsMenuItem struct {
	Name     string
	Command  *commands.DisplayCardsCommand
	PostMenu MenuComponent
}

func NewDisplayCardsMenuItem(name string, command *commands.DisplayCardsCommand, postMenu MenuComponent) *DisplayCardsMenuItem {
	return &DisplayCardsMenuItem{
		Name:     name,
		Command:  command,
		PostMenu: postMenu,
	}
}

func (mi *DisplayCardsMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *DisplayCardsMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))

	///////////////////// To do
	// var userLoggingin models.LoginUserReq
	// userLoggingin.PhoneOrEmail = utils.ReadInput(scanner, "Phone/Email: ")
	// userLoggingin.Password = utils.ReadInput(scanner, "Password: ")
	err := mi.Command.Execute()
	if err != nil {
		fmt.Println(err)
		utils.ReadInput(scanner, "Press any key to continue... ")

		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\tSuccessful Login :)\n")
		utils.ReadInput(scanner, "Press any key to continue... ")
	}
	/////////////////////////
	// TODO: mi.state = ProfileMngmnt state
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *DisplayCardsMenuItem) GetName() string {
	return mi.Name
}
