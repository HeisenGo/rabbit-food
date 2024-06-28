package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/models"
	"client/utils"
	"fmt"
	"strconv"
	"time"
)

type CreateRestaurantMenuItem struct {
	Name     string
	Command  *commands.CreateRestaurantCommand
	PostMenu MenuComponent
}

func NewCreateRestaurantMenuItem(name string, command *commands.CreateRestaurantCommand, postMenu MenuComponent) *CreateRestaurantMenuItem {
	return &CreateRestaurantMenuItem{
		Name:     name,
		Command:  command,
		PostMenu: postMenu,
	}
}

func (mi *CreateRestaurantMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *CreateRestaurantMenuItem) Execute(scanner *bufio.Scanner) {
	defer time.Sleep(time.Second)
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n\n", mi.Name))
	var restaurant models.Restaurant
	var restaurantAddress models.Address
	var coordinates models.Coordinates
	var err error
	restaurant.Name = utils.ReadInput(scanner, "Restaurant Name: ")
	restaurant.Phone = utils.ReadInput(scanner, "Restaurant Phone: ")
	utils.ColoredPrint(constants.Purple, "\n\tAddress:\n")

	restaurantAddress.City = utils.ReadInput(scanner, "City: ")
	restaurantAddress.AddressLine = utils.ReadInput(scanner, "Address Line: ")
	for {
		coordinates.Lat, err = strconv.ParseFloat(utils.ReadInput(scanner, "Coordinates Lat: "), 64)
		if err != nil {
			utils.ColoredPrint(constants.Red, "\n\t", "Float is accepted", "\n")
		}
		if err == nil {
			break
		}
	}

	for {
		coordinates.Lng, err = strconv.ParseFloat(utils.ReadInput(scanner, "Coordinates Lng: "), 64)
		if err != nil {
			utils.ColoredPrint(constants.Red, "\n\t", "Float is accepted", "\n")
		}
		if err == nil {
			break
		}
	}

	restaurantAddress.Coordinates = coordinates
	restaurant.Address = &restaurantAddress
	newRestaurant, err := mi.Command.Execute(restaurant)
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err, "\n")
		utils.ReadInput(scanner, "\n\tPress any key to go back... ")
		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\tRestaurant \"", newRestaurant.Name, "\" Successfully Created!\n")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
	}
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *CreateRestaurantMenuItem) GetName() string {
	return mi.Name
}
