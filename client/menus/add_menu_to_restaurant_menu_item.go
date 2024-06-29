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

type AddMenuToRestaurantMenuItem struct {
	Name                             string
	Command                          *commands.AddMenuToRestaurantCommand
	GetRestaurantsIHaveARoleCommand  *commands.GetRestaurantsIHaveARoleCommand
	GetCategoriesOfRestaurantCommand *commands.GetCategoriesOfRestaurantCommand
	PostMenu                         MenuComponent
}

func NewAddMenuToRestaurantMenuItem(name string, command *commands.AddMenuToRestaurantCommand, getRestaurantsIHaveARoleCommand *commands.GetRestaurantsIHaveARoleCommand, getCategoriesOfRestaurantCommand *commands.GetCategoriesOfRestaurantCommand, postMenu MenuComponent) *AddMenuToRestaurantMenuItem {
	return &AddMenuToRestaurantMenuItem{
		Name:                             name,
		Command:                          command,
		GetRestaurantsIHaveARoleCommand:  getRestaurantsIHaveARoleCommand,
		GetCategoriesOfRestaurantCommand: getCategoriesOfRestaurantCommand,
		PostMenu:                         postMenu,
	}
}

func (mi *AddMenuToRestaurantMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *AddMenuToRestaurantMenuItem) Execute(scanner *bufio.Scanner) {
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

	var restaurantMenuBody tcp.RestaurantMenuBody
	var restaurantRow int
	for {
		input := utils.ReadInput(scanner, "Choose Restaurant Row to Add Menu(Enter q to return): ")

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

	restaurantMenuBody.RestaurantID = restaurants[restaurantRow-1].ID
	// etch categories with getCategoriesCommand
	// separate them with ","
	// add category ids to body
	// entered_category_ids := []uint{}
	// for {
	// 	input := utils.ReadInput(scanner, "\nPlease insert row numbers of your restaurant Categories like: 1, 2, 3  (q to exit):")
	// 	if input == "q" {
	// 		return
	// 	}
	// 	entered_categories := strings.Split(input, ",")
	// 	if len(entered_categories) == 0 {
	// 		continue
	// 	} else {
	// 		ok := true
	// 		for i := 0; i < len(entered_categories); i++ {
	// 			a, err := strconv.Atoi(entered_categories[i])
	// 			if err != nil || a > len(categories) {
	// 				fmt.Println("\n\t Incorrect format")
	// 				ok = false
	// 				break
	// 			}
	// 			entered_category_ids = append(entered_category_ids, categories[a-1].ID)
	// 		}
	// 		if ok {
	// 			break
	// 		}
	// 	}
	// }
	// if len(entered_category_ids) == 0 {
	// 	return
	// }
	//restaurantCategoryBody.Category_ids = entered_category_ids

	input := utils.ReadInput(scanner, "\nPlease insert Menu Name: ")
	restaurantMenuBody.Name = input
	err = mi.Command.Execute(restaurantMenuBody)
	if err != nil {
		utils.ColoredPrint(constants.Red, "\n\t", err, "\n")
		utils.ReadInput(scanner, "\n\tPress any key to go back... ")
		return
	} else {
		utils.ColoredPrint(constants.Green, "\n\tCategories Successfully added :)\n")
		utils.ReadInput(scanner, "\n\tPress any key to continue... ")
	}
	if mi.PostMenu != nil {
		mi.PostMenu.Execute(scanner)
		return
	}
}

func (mi *AddMenuToRestaurantMenuItem) GetName() string {
	return mi.Name
}
