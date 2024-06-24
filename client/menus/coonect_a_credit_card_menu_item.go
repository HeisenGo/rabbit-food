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

type ConnectACreditCardMenuItem struct {
	Name     string
	Command  *commands.ConnectACreditCardCommand
	PostMenu MenuComponent
}

func NewConnectACreditCardMenuItem(name string, command *commands.ConnectACreditCardCommand, postMenu MenuComponent) *ConnectACreditCardMenuItem {
	return &ConnectACreditCardMenuItem{
		Name:     name,
		Command:  command,
		PostMenu: postMenu,
	}
}

func (mi *ConnectACreditCardMenuItem) Display() {
	fmt.Println(mi.Name)
}

// func (mi *LoginMenuItem) Execute(scanner *bufio.Scanner) {
// 	defer time.Sleep(time.Second)
// 	utils.ClearScreen()
// 	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))
// }

func (mi *ConnectACreditCardMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))
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
	// TODO: mi.state = ProfileMngmnt state
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *ConnectACreditCardMenuItem) GetName() string {
	return mi.Name
}
