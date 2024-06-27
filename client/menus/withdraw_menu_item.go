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

type WithdrawMenuItem struct {
	Name                string
	WithdrawCommand     *commands.WithdrawCommand
	DisplayCardsCommand *commands.DisplayCardsCommand
	PostMenu            MenuComponent
}

func NewWithdrawMenuItem(name string, withdrawCommand *commands.WithdrawCommand, DisplayCardsCommand *commands.DisplayCardsCommand, postMenu MenuComponent) *WithdrawMenuItem {
	return &WithdrawMenuItem{
		Name:                name,
		WithdrawCommand:     withdrawCommand,
		DisplayCardsCommand: DisplayCardsCommand,
		PostMenu:            postMenu,
	}
}

func (mi *WithdrawMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *WithdrawMenuItem) Execute(scanner *bufio.Scanner) {
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n\n", mi.Name))
	cards, err := mi.DisplayCardsCommand.Execute()
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

	var withdrawBody tcp.WithdrawBody
	cardRow, err := strconv.Atoi(utils.ReadInput(scanner, "Choose Card Row to Deposit: "))
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", errors.ErrDataType.Message)
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}
	if cardRow > len(cards) {
		utils.ColoredPrint(constants.Red, "\tInvalid number")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}
	withdrawBody.Number = cards[cardRow-1].Number
	amount, err := strconv.Atoi(utils.ReadInput(scanner, "Amount: "))
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", errors.ErrDataType.Message)
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}
	if amount <= 0 {
		utils.ColoredPrint(constants.Red, "\n\t", "Entered Amount should be positive")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}
	withdrawBody.Amount = uint(amount)

	err = mi.WithdrawCommand.Execute(&withdrawBody)
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

func (mi *WithdrawMenuItem) GetName() string {
	return mi.Name
}
