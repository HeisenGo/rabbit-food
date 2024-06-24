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

type LoginMenuItem struct {
	Name     string
	Command  *commands.LoginCommand
	PostMenu MenuComponent
}

func NewLoginMenuItem(name string, command *commands.LoginCommand, postMenu MenuComponent) *LoginMenuItem {
	return &LoginMenuItem{
		Name:     name,
		Command:  command,
		PostMenu: postMenu,
	}
}

func (mi *LoginMenuItem) Display() {
	fmt.Println(mi.Name)
}

// func (mi *LoginMenuItem) Execute(scanner *bufio.Scanner) {
// 	defer time.Sleep(time.Second)
// 	utils.ClearScreen()
// 	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))
// }

func (mi *LoginMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))
	var userLoginData tcp.LoginBody
	userLoginData.PhoneOrEmail = utils.ReadInput(scanner, "Phone/Email: ")
	userLoginData.Password = utils.ReadInput(scanner, "Password: ")
	err := mi.Command.Execute(&userLoginData)
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

func (mi *LoginMenuItem) GetName() string {
	return mi.Name
}
