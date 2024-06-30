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

type AddMotorToRestaurantMenuItem struct {
	Name              string
	Command           *commands.AddMotorToRestaurantCommand
	GetRestaurantsIOwn *commands.GetRestaurantsIOwnCommand
	PostMenu          MenuComponent
}

func NewAddMotorToRestaurantMenuItem(name string, command *commands.AddMotorToRestaurantCommand, getRestaurantsIHaveCommand *commands.GetRestaurantsIOwnCommand, postMenu MenuComponent) *AddMotorToRestaurantMenuItem {
	return &AddMotorToRestaurantMenuItem{
		Name:                             name,
		Command:                          command,
		GetRestaurantsIOwn:  getRestaurantsIHaveCommand,
		PostMenu:                         postMenu,
	}
}

func (mi *AddMotorToRestaurantMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *AddMotorToRestaurantMenuItem) Execute(scanner *bufio.Scanner) {
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

	utils.ColoredPrint(constants.Green, "\n\tRestaurants You can add Motor to: \n")
	functions.DisplayRestaurantsWithAddress(restaurants)
	fmt.Println("\n\n\t")

	var addMotorReqBody tcp.RestaurantMotorReqBody
	var restaurantRow int
	for {
		input := utils.ReadInput(scanner, "Choose Restaurant Row to Add Motor(Enter q to return): ")

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

	addMotorReqBody.RestaurantID = restaurants[restaurantRow-1].ID

	input := utils.ReadInput(scanner, "\nPlease insert Name of Motor: ")
	addMotorReqBody.Name = input
	input = utils.ReadInput(scanner, "\nPlease insert Speed of Motor(int): ")
	addMotorReqBody.Speed, err = strconv.Atoi(input)
	if err!=nil{
		utils.ColoredPrint(constants.Red, "\n\t", errors.ErrDataType, "\n")
		utils.ReadInput(scanner, "\n\tPress any key to go back... ")
		return
	}
	err = mi.Command.Execute(addMotorReqBody)
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err, "\n")
		utils.ReadInput(scanner, "\n\tPress any key to go back... ")
		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\tMotor Successfully added :)\n")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
	}
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *AddMotorToRestaurantMenuItem) GetName() string {
	return mi.Name
}
