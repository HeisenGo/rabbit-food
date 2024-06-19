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

type WalletMenuItem struct {
	Name     string
	Command  *commands.GetWalletCommand
	PostMenu MenuComponent
}

func NewWalletMenuItem(name string, command *commands.GetWalletCommand, postMenu MenuComponent) *WalletMenuItem {
	return &WalletMenuItem{
		Name:     name,
		Command:  command,
		PostMenu: postMenu,
	}
}

func (mi *WalletMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *WalletMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))
	var wallet models.GetWalletReq
	// TODO: should get from context
	wallet.ID = 4
	receivedWallet, err := mi.Command.Execute(&wallet)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\n\tYour Wallet Balance: %d\n\n", receivedWallet.Balance)
	utils.ReadInput(scanner, "Press any keys to back: ")
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
	}
}
