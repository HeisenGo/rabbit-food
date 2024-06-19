package menus

import (
	"bufio"
	"client/utils"
	"fmt"
	"strconv"
)

type Menu struct {
	Name      string
	MenuItems []MenuComponent
}

func (m *Menu) Display() {
	fmt.Println("\n---", m.Name, "---")
	for i, menuItem := range m.MenuItems {
		fmt.Printf("%d. ", i+1)
		menuItem.Display()
	}
	fmt.Printf("%d. Return to previous menu\n", len(m.MenuItems)+1)
}

func (m *Menu) Execute(scanner *bufio.Scanner) {
	utils.ClearScreen()
	for {
		m.Display()
		fmt.Print("Enter choice: ")
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
		m.MenuItems[choice-1].Execute(scanner)
	}
}

func (m *Menu) add(mc MenuComponent) {
	m.MenuItems = append(m.MenuItems, mc)
}
