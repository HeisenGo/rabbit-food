package menus

import (
	"bufio"
	"client/constants"
	"client/utils"
	"fmt"
	"strconv"
)

type Menu struct {
	Name      string
	MenuItems []MenuComponent
	// TODO: nextState  menuState
	// TODO: state  menuState
}

func (m *Menu) Display() {
	utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", m.Name))
	for i, menuItem := range m.MenuItems {
		fmt.Printf("\t%d. ", i+1)
		menuItem.Display()
	}
	fmt.Printf("\t%d. Return to previous menu\n", len(m.MenuItems)+1)
}

func (m *Menu) Execute(scanner *bufio.Scanner) {
	// TODO: setState -> m.setState(m.state.next)
	for {
		utils.ClearScreen()
		m.Display()
		utils.ColoredPrint(constants.Green, "\t Enter Choice: ")
		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(m.MenuItems)+1 {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		if choice == len(m.MenuItems)+1 {
			utils.ClearScreen()
			return // Return to previous menu
		}
		// TODO: mi.state.execute
		m.MenuItems[choice-1].Execute(scanner)
	}
}

// TODO: Should be use and export
func (m *Menu) add(mc MenuComponent) {
	m.MenuItems = append(m.MenuItems, mc)
}
