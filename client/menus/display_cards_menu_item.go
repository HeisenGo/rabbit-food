package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/utils"
	"fmt"
	"time"
)

type DisplayCardsMenuItem struct {
	Name     string
	Command  *commands.DisplayCardsCommand
	PostMenu MenuComponent
}

func NewDisplayCardsMenuItem(name string, command *commands.DisplayCardsCommand, postMenu MenuComponent) *DisplayCardsMenuItem {
	return &DisplayCardsMenuItem{
		Name:     name,
		Command:  command,
		PostMenu: postMenu,
	}
}

func (mi *DisplayCardsMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *DisplayCardsMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n\n", mi.Name))

	cards, err := mi.Command.Execute()
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err)
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\tCards: \n")
		for i, card := range cards {
			fmt.Printf("\n\t %v. %v", i+1, card.Number)
		}
		fmt.Println("")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
	}
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *DisplayCardsMenuItem) GetName() string {
	return mi.Name
}
