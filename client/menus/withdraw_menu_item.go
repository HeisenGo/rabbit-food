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
	if len(cards) == 0 {
		utils.ColoredPrint(constants.Red, "\n\t", "You have added no cards!")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}
	utils.ColoredPrint(constants.Green, "\n\tCards: \n")
	for i, card := range cards {
		fmt.Printf("\n\t %v. %v", i+1, utils.SeparateByFour(card.Number))
	}
	fmt.Println("\n\n\t")

	var withdrawBody tcp.WithdrawBody
	var cardRow int
	var amount int
	for {
		input := utils.ReadInput(scanner, "Choose Card Row to Deposit(Enter q to return): ")

		if input == "q" {
			return
		}
		cardRow, err = strconv.Atoi(input)

		if err != nil {
			utils.ColoredPrint(constants.Red, "\t", errors.ErrDataType.Message, "\n")
			continue
		}
		if cardRow > len(cards) {
			utils.ColoredPrint(constants.Red, "\tInvalid number\n")
			continue
		}
		if cardRow <= len(cards) {
			break
		}
	}
	withdrawBody.Number = cards[cardRow-1].Number

	for {
		input := utils.ReadInput(scanner, "Amount (q to exit process): ")

		if input == "q" {
			return
		}
		amount, err = strconv.Atoi(input)

		if err != nil {
			utils.ColoredPrint(constants.Red, "\t", errors.ErrDataType.Message, "\n")
			continue
		}
		if amount <= 0 {
			utils.ColoredPrint(constants.Red, "\n\t", "Entered Amount should be positive\n")
			continue
		}
		if amount >= 0 {
			break
		}
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
