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

type AddOperatorToRestaurantMenuItem struct {
	Name               string
	Command            *commands.AddOperatorToRestaurantCommand
	GetRestaurantsIOwn *commands.GetRestaurantsIOwnCommand
	PostMenu           MenuComponent
}

func NewAddOperatorToRestaurantMenuItem(name string, command *commands.AddOperatorToRestaurantCommand, getRestaurantsIHaveCommand *commands.GetRestaurantsIOwnCommand, postMenu MenuComponent) *AddOperatorToRestaurantMenuItem {
	return &AddOperatorToRestaurantMenuItem{
		Name:               name,
		Command:            command,
		GetRestaurantsIOwn: getRestaurantsIHaveCommand,
		PostMenu:           postMenu,
	}
}

func (mi *AddOperatorToRestaurantMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *AddOperatorToRestaurantMenuItem) Execute(scanner *bufio.Scanner) {
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n\n", mi.Name))
	restaurants, err := mi.GetRestaurantsIOwn.Execute()

	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err)
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}
	if len(restaurants) == 0 {
		utils.ColoredPrint(constants.Red, "\n\t", "You have no restaurants")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}

	utils.ColoredPrint(constants.Green, "\n\tRestaurants You can add Operator to: \n")
	functions.DisplayRestaurantsWithAddress(restaurants)
	fmt.Println("\n\n\t")

	var addOperatorReqBody tcp.RestaurantAddOperatorReqBody
	var restaurantRow int
	for {
		input := utils.ReadInput(scanner, "Choose Restaurant Row to Add Operator(Enter q to return): ")

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

	addOperatorReqBody.RestaurantID = restaurants[restaurantRow-1].ID

	input := utils.ReadInput(scanner, "\nPlease insert Registered Operator Phone/Email: ")
	addOperatorReqBody.PhoneOrEmail = input

	err = mi.Command.Execute(addOperatorReqBody)
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err, "\n")
		utils.ReadInput(scanner, "\n\tPress any key to go back... ")
		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\tOperator Successfully added :)\n")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
	}
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *AddOperatorToRestaurantMenuItem) GetName() string {
	return mi.Name
}
