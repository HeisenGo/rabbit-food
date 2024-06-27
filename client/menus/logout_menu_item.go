package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/utils"
	"fmt"
	"time"
)

type LogoutMenuItem struct {
	Name     string
	Command  *commands.LogoutCommand
	PostMenu MenuComponent
}

func NewLogoutMenuItem(name string, command *commands.LogoutCommand, postMenu MenuComponent) *LogoutMenuItem {
	return &LogoutMenuItem{
		Name:     name,
		Command:  command,
		PostMenu: postMenu,
	}
}

func (mi *LogoutMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *LogoutMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n\n", mi.Name))
	areYouSureToLogout := utils.ReadInput(scanner, "Are you sure to logout?(y/n): ")
	if areYouSureToLogout == "y" {
		mi.PostMenu.Execute(scanner)
	}
}

func (mi *LogoutMenuItem) GetName() string {
	return mi.Name
}
