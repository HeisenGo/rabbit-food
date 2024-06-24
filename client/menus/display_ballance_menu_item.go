package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/models"
	"client/utils"
	"fmt"
	"time"
)

type DisplayBallanceMenuItem struct {
	Name     string
	Command  *commands.DisplayBallanceCommand
	PostMenu MenuComponent
}

func NewDisplayBalanceMenuItem(name string, command *commands.DisplayBallanceCommand, postMenu MenuComponent) *DisplayBallanceMenuItem {
	return &DisplayBallanceMenuItem{
		Name:     name,
		Command:  command,
		PostMenu: postMenu,
	}
}

func (mi *DisplayBallanceMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *DisplayBallanceMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))

	///////////////////// To do
	var userLoggingin models.LoginUserReq
	userLoggingin.PhoneOrEmail = utils.ReadInput(scanner, "Phone/Email: ")
	userLoggingin.Password = utils.ReadInput(scanner, "Password: ")
	err := mi.Command.Execute(&userLoggingin)
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

func (mi *DisplayBallanceMenuItem) GetName() string {
	return mi.Name
}
