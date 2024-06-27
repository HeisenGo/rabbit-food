package menus

import (
	"bufio"
	"client/commands"
	"client/constants"
	"client/models"
	"client/utils"
	"fmt"
	"strconv"
)

type ProfileMenuItem struct {
	Name              string
	DisplayProfileCmd *commands.DisplayProfileCommand
	EditProfileCmd    *commands.EditProfileCommand
	PostMenu          MenuComponent
}

func NewProfileMenuItem(name string, displayProfileCmd *commands.DisplayProfileCommand, editProfileCmd *commands.EditProfileCommand, postMenu MenuComponent) *ProfileMenuItem {
	return &ProfileMenuItem{
		Name:              name,
		DisplayProfileCmd: displayProfileCmd,
		EditProfileCmd:    editProfileCmd,
		PostMenu:          postMenu,
	}
}

func (mi *ProfileMenuItem) Display() {
	fmt.Println(mi.Name)
}

func (mi *ProfileMenuItem) Execute(scanner *bufio.Scanner) {
	for {
		utils.ClearScreen()
		utils.ColoredPrint(constants.Blue, fmt.Sprintf("[------------ %s ------------] \n", mi.Name))
		fmt.Println("1. Display Profile Information")
		fmt.Println("2. Edit User Information")
		fmt.Println("3. Back to Main Menu")
		utils.ColoredPrint(constants.Green, "\t Enter Choice: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			mi.displayProfile(scanner)
		case "2":
			mi.editProfile(scanner)
		case "3":
			utils.ClearScreen()
			return // Return to previous menu
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func (mi *ProfileMenuItem) displayProfile(scanner *bufio.Scanner) {
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, "[--- Display Profile Information ---] \n")
	fmt.Print("Enter user ID: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid user ID")
		utils.ReadInput(scanner, "Press any key to continue... ")
		return
	}
	err = mi.DisplayProfileCmd.Execute(uint(id))
	if err != nil {
		fmt.Println("Error displaying profile:", err)
	} else {
		utils.ColoredPrint(constants.Green, "\n\tProfile displayed successfully!\n")
	}
	utils.ReadInput(scanner, "Press any key to continue... ")
}

func (mi *ProfileMenuItem) editProfile(scanner *bufio.Scanner) {
	utils.ClearScreen()
	utils.ColoredPrint(constants.Blue, "[--- Edit Profile Information ---] \n")
	var user models.User
	fmt.Print("Enter user ID: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid user ID")
		utils.ReadInput(scanner, "Press any key to continue... ")
		return
	}
	user.ID = uint(id)
	user.FirstName = utils.ReadInput(scanner, "First Name: ")
	user.LastName = utils.ReadInput(scanner, "Last Name: ")
	user.Phone = utils.ReadInput(scanner, "Phone: ")
	email := utils.ReadInput(scanner, "Email: ")
	user.Email = email

	err = mi.EditProfileCmd.Execute(&user)
	if err != nil {
		fmt.Println("Error editing profile:", err)
	} else {
		utils.ColoredPrint(constants.Green, "\n\tProfile updated successfully!\n")
	}
	utils.ReadInput(scanner, "Press any key to continue... ")
}

func (mi *ProfileMenuItem) GetName() string {
	return mi.Name
}
