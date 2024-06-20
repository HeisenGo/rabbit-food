package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
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

func (mi *LoginMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))
}

func (mi *LoginMenuItem) GetName() string {
	return mi.Name
}
