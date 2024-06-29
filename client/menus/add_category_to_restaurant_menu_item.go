package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/protocol/tcp"
	"client/utils"
	"fmt"
	"time"
)

type AddCategoryToRestaurantMenuItem struct {
	Name     string
	Command  *commands.AddCategoryToRestaurantCommand
	GetRestaurantsIHaveARoleCommand   commands.GetRestaurantsIHaveARoleCommand
	PostMenu MenuComponent
}

func NewAddCategoryToRestaurantMenuItem(name string, command *commands.AddCategoryToRestaurantCommand, getRestaurantsIHaveARoleCommand   commands.GetRestaurantsIHaveARoleCommand, postMenu MenuComponent) *AddCategoryToRestaurantMenuItem {
	return &AddCategoryToRestaurantMenuItem{
		Name:     name,
		Command:  command,
		GetRestaurantsIHaveARoleCommand: getRestaurantsIHaveARoleCommand,
		PostMenu: postMenu,
	}
}

func (mi *AddCategoryToRestaurantMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *AddCategoryToRestaurantMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n\n", mi.Name))
	var addCardData tcp.AddCardBody
	addCardData.CardNumber = utils.ReadInput(scanner, "Card Number: ")
	newCard, err := mi.Command.Execute(&addCardData)
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err, "\n")
		utils.ReadInput(scanner, "\n\tPress any key to go back... ")
		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\tCard "+newCard.Number+" Successfully added :)\n")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
	}
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *AddCategoryToRestaurantMenuItem) GetName() string {
	return mi.Name
}
