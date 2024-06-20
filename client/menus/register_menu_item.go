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

type RegisterMenuItem struct {
	Name     string
	Command  *commands.RegisterCommand
	PostMenu MenuComponent
}

func NewRegisterMenuItem(name string, command *commands.RegisterCommand, postMenu MenuComponent) *RegisterMenuItem {
	return &RegisterMenuItem{
		Name:     name,
		Command:  command,
		PostMenu: postMenu,
	}
}

func (mi *RegisterMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *RegisterMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))
	var user models.User
	user.Phone = utils.ReadInput(scanner, "Phone: ")
	user.Email = utils.ReadInput(scanner, "Email: ")
	user.Password = utils.ReadInput(scanner, "Password: ")
	err := mi.Command.Execute(&user)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\tSuccessful Registration!\n")
		utils.ReadInput(scanner, "Press any key to continue... ")
	}
	// TODO: mi.state = ProfileMngmnt state
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *RegisterMenuItem) GetName() string {
	return mi.Name
}
