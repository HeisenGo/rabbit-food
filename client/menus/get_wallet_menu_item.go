package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/utils"
	"fmt"
	"strconv"
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
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))
	wallet, err := mi.Command.Execute()
	if err != nil {
		utils.ColoredPrint(constants.Red, err)
		utils.ReadInput(scanner, "\nPress any key to continue... ")

		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\tYour Ballance is: ", strconv.Itoa(wallet.Balance))
		utils.ReadInput(scanner, "Press any key to continue... ")
	}
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *GetWalletMenuItem) GetName() string {
	return mi.Name
}
