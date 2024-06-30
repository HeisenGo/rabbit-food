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

type AddItemToMenuMenuItem struct {
	Name                            string
	Command                         *commands.AddItemToMenuCommand
	GetRestaurantsIHaveARoleCommand *commands.GetRestaurantsIHaveARoleCommand
	GetRestaurantMenus              *commands.GetMenusOfRestaurantCommand
	PostMenu                        MenuComponent
}

func NewAddItemToMenuMenuItem(name string, command *commands.AddItemToMenuCommand, getRestaurantsIHaveARoleCommand *commands.GetRestaurantsIHaveARoleCommand, getMenusCommand *commands.GetMenusOfRestaurantCommand, postMenu MenuComponent) *AddItemToMenuMenuItem {
	return &AddItemToMenuMenuItem{
		Name:                            name,
		Command:                         command,
		GetRestaurantsIHaveARoleCommand: getRestaurantsIHaveARoleCommand,
		GetRestaurantMenus:              getMenusCommand,
		PostMenu:                        postMenu,
	}
}

func (mi *AddItemToMenuMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *AddItemToMenuMenuItem) Execute(scanner *bufio.Scanner) {
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

	menus, err := mi.GetRestaurantMenus.Execute(restaurants[restaurantRow-1].ID)
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err, "\n")
		utils.ReadInput(scanner, "\n\tPress any key to go back... ")
		return
	}
	if len(menus) == 0 {
		utils.ColoredPrint(constants.Red, "\n\t", "You need to add Menu first", "\n")
		utils.ReadInput(scanner, "\n\tPress any key to go back... ")
		return
	}

	functions.DisplayMenus(menus)

	var menuRow int
	for {
		input := utils.ReadInput(scanner, "Choose Menu Row to Add Item(Enter q to return): ")

		if input == "q" {
			return
		}
		menuRow, err = strconv.Atoi(input)

		if err != nil {
			utils.ColoredPrint(constants.Red, "\t", errors.ErrDataType.Message, "\n")
			continue
		}
		if menuRow > len(menus) {
			utils.ColoredPrint(constants.Red, "\tInvalid number\n")
			continue
		}
		if menuRow <= len(menus) {
			break
		}
	}

	var requestBody tcp.AddItemToMenuReqBody
	requestBody.MenuID = menus[menuRow-1].ID
	requestBody.Name = utils.ReadInput(scanner, "Item Name: ")
	var price int
	for {
		price, err = strconv.Atoi(utils.ReadInput(scanner, "Item Price: "))
		if err != nil {
			utils.ColoredPrint(constants.Red, "\n\t", errors.ErrDataType, "\n")
		} else {
			break
		}
	}
	requestBody.Price = uint(price)

	var preparationMinutes int
	for {
		preparationMinutes, err = strconv.Atoi(utils.ReadInput(scanner, "Item Time to Prepare in minutes: "))
		if err != nil {
			utils.ColoredPrint(constants.Red, "\n\t", errors.ErrDataType, "\n")
		} else {
			break
		}
	}
	requestBody.PreparationMinutes = uint(preparationMinutes)

	var penalty float64
	for {
		penalty, err = strconv.ParseFloat(utils.ReadInput(scanner, "Item Cancellation Penalty Percentage: "), 64)
		if err != nil {
			utils.ColoredPrint(constants.Red, "\n\t", errors.ErrDataType, "\n")
		} else {
			break
		}
	}
	requestBody.CancellationPenaltyPercentage = uint(penalty)

	err = mi.Command.Execute(requestBody)
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err, "\n")
		utils.ReadInput(scanner, "\n\tPress any key to go back... ")
		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\tItem Successfully added :)\n")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
	}
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *AddItemToMenuMenuItem) GetName() string {
	return mi.Name
}
