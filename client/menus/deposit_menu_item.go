package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/errors"
	"client/protocol/tcp"
	"client/utils"
	"fmt"
	"strconv"
)

type DepositMenuItem struct {
	Name                  string
	DepositCommand        *commands.DepositCommand
	GetWalletCardsCommand *commands.DisplayCardsCommand
	PostMenu              MenuComponent
}

func NewDepositMenuItem(name string, depositCommand *commands.DepositCommand, getWalletCardsCommand *commands.DisplayCardsCommand, postMenu MenuComponent) *DepositMenuItem {
	return &DepositMenuItem{
		Name:                  name,
		DepositCommand:        depositCommand,
		GetWalletCardsCommand: getWalletCardsCommand,
		PostMenu:              postMenu,
	}
}

func (mi *DepositMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *DepositMenuItem) Execute(scanner *bufio.Scanner) {
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))
	cards, err := mi.GetWalletCardsCommand.Execute()
	if err != nil {
		utils.ColoredPrint(constants.Red, err)
		utils.ReadInput(scanner, "\nPress any key to continue... ")
		return
	}
	utils.ColoredPrint(constants.Green, "\n\tCards: \n")
	for i, c := range cards {
		utils.ColoredPrint(constants.Purple, fmt.Sprintf("\n\t%d- %v", i+1, c.Number))
	}
	fmt.Println("\n\tPlease Choose Which card and How Much to Deposit")

	var depositBody tcp.DepositBody
	cardRow, err := strconv.Atoi(utils.ReadInput(scanner, "Card Row: "))
	if err != nil {
		utils.ColoredPrint(constants.Red, errors.ErrDataType.Message)
		utils.ReadInput(scanner, "\nPress any key to continue... ")
		return
	}
	depositBody.Number = cards[cardRow-1].Number
	amount, err := strconv.Atoi(utils.ReadInput(scanner, "Amount: "))
	if err != nil {
		utils.ColoredPrint(constants.Red, errors.ErrDataType.Message)
		utils.ReadInput(scanner, "\nPress any key to continue... ")
		return
	}
	depositBody.Amount = uint(amount)

	err = mi.DepositCommand.Execute(&depositBody)
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err, "\n")
		utils.ReadInput(scanner, "Press any key to go back... ")
		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\tSuccessful Deposit :)\n")
		utils.ReadInput(scanner, "Press any key to continue... ")
	}
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *DepositMenuItem) GetName() string {
	return mi.Name
}
