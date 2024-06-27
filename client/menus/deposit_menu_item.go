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
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n\n", mi.Name))
	cards, err := mi.GetWalletCardsCommand.Execute()
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err)
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}
	utils.ColoredPrint(constants.Green, "\n\tCards: \n")
	for i, card := range cards {
		fmt.Printf("\n\t %v. %v", i+1, utils.SeparateByFour(card.Number))
	}
	fmt.Println("\n\n\t")

	var depositBody tcp.DepositBody
	cardRow, err := strconv.Atoi(utils.ReadInput(scanner, "Choose Card Row to Deposit: "))

	if err != nil {
		utils.ColoredPrint(constants.Red, "\t", errors.ErrDataType.Message)
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}
	if cardRow > len(cards) {
		utils.ColoredPrint(constants.Red, "\tInvalid number")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}
	depositBody.Number = cards[cardRow-1].Number
	amount, err := strconv.Atoi(utils.ReadInput(scanner, "Amount: "))
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", errors.ErrDataType.Message)
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}
	depositBody.Amount = uint(amount)

	err = mi.DepositCommand.Execute(&depositBody)
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err, "\n")
		utils.ReadInput(scanner, "\n\tPress any key to go back... ")
		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\tSuccessful Deposit :)\n")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
	}
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *DepositMenuItem) GetName() string {
	return mi.Name
}
