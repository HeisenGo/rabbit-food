package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/menus/functions"
	"client/utils"
	"fmt"
	"time"
)

type GetRestaurantsIHaveARoleInitMenuItem struct {
	Name     string
	Command  *commands.GetRestaurantsIHaveARoleCommand
	PostMenu MenuComponent
}

func NewGetRestaurantsIHaveARoleInitMenuItem(name string, command *commands.GetRestaurantsIHaveARoleCommand, postMenu MenuComponent) *GetRestaurantsIHaveARoleInitMenuItem {
	return &GetRestaurantsIHaveARoleInitMenuItem{
		Name:     name,
		Command:  command,
		PostMenu: postMenu,
	}
}

func (mi *GetRestaurantsIHaveARoleInitMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *GetRestaurantsIHaveARoleInitMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n\n", mi.Name))

	restaurants, err := mi.Command.Execute()
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err)
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}
	utils.ColoredPrint(constants.Green, "\n\tRestaurants: \n\n")
	if len(restaurants) == 0 {
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
		return
	}
	functions.DisplayRestaurantsWithAddress(restaurants)
	fmt.Println("")
	utils.ReadInput(scanner, "\n\tPress any key to continue... ")

	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *GetRestaurantsIHaveARoleInitMenuItem) GetName() string {
	return mi.Name
}
