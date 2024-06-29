package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/utils"
	"fmt"
	"time"
)

type GetWalletMenuItem struct {
	Name     string
	Command  *commands.GetWalletCommand
	PostMenu MenuComponent
}

func NewGetWalletMenuItem(name string, command *commands.GetWalletCommand, postMenu MenuComponent) *GetWalletMenuItem {
	return &GetWalletMenuItem{
		Name:     name,
		Command:  command,
		PostMenu: postMenu,
	}
}

func (mi *GetWalletMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *GetWalletMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n\n", mi.Name))
	wallet, err := mi.Command.Execute()
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err)
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")

		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\n\tYour Ballance is: ", wallet.Balance, "\n")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
	}
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *GetWalletMenuItem) GetName() string {
	return mi.Name
}
