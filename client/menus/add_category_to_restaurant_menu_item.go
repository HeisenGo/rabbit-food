package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/errors"
	"client/menus/functions"
	"client/models"
	"client/protocol/tcp"
	"client/utils"
	"fmt"
	"strconv"
	"strings"
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
	categories := make([]models.RestaurantCategory, 2)
	categories[0] = models.RestaurantCategory{
		ID:   1,
		Name: "Irani",
	}
	categories[1] = models.RestaurantCategory{
		ID:   2,
		Name: "Fast Food",
	}

	fmt.Println("\n\tCategories:")

	for _, c := range categories{
		fmt.Println("\t",c.Name)
	}
	entered_category_ids := []uint{}
	for {
		input := utils.ReadInput(scanner, "\nPlease insert row numbers of your restaurant Categories like: 1, 2, 3  (q to exit):")
		if input == "q" {
			return
		}
		entered_categories := strings.Split(input, ",")
		if len(entered_categories) == 0 {
			continue
		} else {
			ok := true
			for i := 0; i < len(entered_categories); i++ {
				a, err := strconv.Atoi(entered_categories[i])
				if err != nil || a > len(categories) {
					fmt.Println("\n\t Incorrect format")
					ok = false
					break
				}
				entered_category_ids = append(entered_category_ids, categories[a-1].ID)
			}
			if ok {
				break
			}
		}
	}
	if len(entered_category_ids) == 0 {
		return
	}
	restaurantCategoryBody.Category_ids = entered_category_ids

	err = mi.Command.Execute(restaurantCategoryBody)
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

func (mi *AddCategoryToRestaurantMenuItem) GetName() string {
	return mi.Name
}
