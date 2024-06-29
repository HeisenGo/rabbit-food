package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/errors"
	"client/menus/functions"
	"client/protocol/tcp"
	"client/utils"
	"fmt"
	"strconv"
)

type AddCategoryToRestaurantMenuItem struct {
	Name                            string
	Command                         *commands.AddCategoryToRestaurantCommand
	GetRestaurantsIHaveARoleCommand *commands.GetRestaurantsIHaveARoleCommand
	PostMenu                        MenuComponent
}

func NewAddCategoryToRestaurantMenuItem(name string, command *commands.AddCategoryToRestaurantCommand, getRestaurantsIHaveARoleCommand *commands.GetRestaurantsIHaveARoleCommand, postMenu MenuComponent) *AddCategoryToRestaurantMenuItem {
	return &AddCategoryToRestaurantMenuItem{
		Name:                            name,
		Command:                         command,
		GetRestaurantsIHaveARoleCommand: getRestaurantsIHaveARoleCommand,
		PostMenu:                        postMenu,
	}
}

func (mi *AddCategoryToRestaurantMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *AddCategoryToRestaurantMenuItem) Execute(scanner *bufio.Scanner) {
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n\n", mi.Name))
	restaurants, err := mi.GetRestaurantsIHaveARoleCommand.Execute()

	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err)
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}
	if len(restaurants) == 0 {
		utils.ColoredPrint(constants.Red, "\n\t", "You have role in any restaurants")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}

	utils.ColoredPrint(constants.Green, "\n\tRestaurants You can add Category: \n")
	functions.DisplayRestaurantsWithAddress(restaurants)
	fmt.Println("\n\n\t")

	var restaurantCategoryBody tcp.RestaurantCategoryBody
	var restaurantRow int
	for {
		input := utils.ReadInput(scanner, "Choose Restaurant Row to Add Category(Enter q to return): ")

		if input == "q" {
			return
		}
		restaurantRow, err = strconv.Atoi(input)

		if err != nil {
			utils.ColoredPrint(constants.Red, "\t", errors.ErrDataType.Message, "\n")
			continue
		}
		if restaurantRow > len(restaurants) {
			utils.ColoredPrint(constants.Red, "\tInvalid number\n")
			continue
		}
		if restaurantRow <= len(restaurants) {
			break
		}
	}

	restaurantCategoryBody.RestaurantID = restaurants[restaurantRow-1].ID
	// etch categories with getCategoriesCommand
	// separate them with "," 
	// add category ids to body
	//categories:=
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
