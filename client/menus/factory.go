package menus

import (
	"client/commands"
	"client/services"
)

var (
	mainMenu         *Menu
	registerMenuItem MenuComponent
	loginMenuItem    MenuComponent
	//////       Home Menu ///////////////
	homeMenu                 *Menu // TODO: bayad be interface vabaste bashe ye fekri bokon
	profileManagementMenu    *Menu
	walletManagementMenu     *Menu
	foodOrderMenu            MenuComponent
	restaurantManagementMenu MenuComponent
	logoutMenuItem           MenuComponent
	//////       Profile ManagementMeu /////
	///// Info to display and last updated address to display
	editInfoMenu    *Menu
	manageAddresses *Menu
	editAddress     MenuComponent
	deleteAddress   MenuComponent
	addNewAddress   MenuComponent

	/// wallet management menu ////
	getBallanceMenuItem   MenuComponent
	displayCardsMenuItem  MenuComponent
	AddCreditCardMenuItem MenuComponent
	depositMenuItem       MenuComponent
	withDrawMenuItem      MenuComponent
)

func getDisplayCardsMenuItem(service services.Service) MenuComponent {
	if displayCardsMenuItem == nil {
		displayCardsMenuItem = NewDisplayCardsMenuItem("Cards", commands.NewDisplayCardsCommand(service), getWalletManagementMenu(service))
	}
	return displayCardsMenuItem
}

func getGetBalanceMenuItem(service services.Service) MenuComponent {
	if getBallanceMenuItem == nil {
		getBallanceMenuItem = NewGetWalletMenuItem("My Ballance", commands.NewGetWalletCommand(service), getWalletManagementMenu(service))
	}
	return getBallanceMenuItem
}

func getAddCardMenuItem(service services.Service) MenuComponent {
	if AddCreditCardMenuItem == nil {
		AddCreditCardMenuItem = NewAddCardMenuItem("Connect A New Credit Card", commands.NewAddCardCommand(service), getWalletManagementMenu(service))
	}
	return AddCreditCardMenuItem
}

func getDepositMenuItem(service services.Service) MenuComponent {
	if depositMenuItem == nil {
		depositMenuItem = NewDepositMenuItem("Deposit", commands.NewDepositCommand(service), commands.NewDisplayCardsCommand(service), getWalletManagementMenu(service))
	}
	return depositMenuItem
}

func getWithDrawMenu(service services.Service) MenuComponent {
	if withDrawMenuItem == nil {
		withDrawMenuItem = NewWithdrawMenuItem("Withdraw", commands.NewWithdrawCommand(service), commands.NewDisplayCardsCommand(service), getWalletManagementMenu(service))
	}
	return withDrawMenuItem
}

func getWalletManagementMenu(service services.Service) MenuComponent {
	if walletManagementMenu == nil {
		walletManagementMenu = NewMenu("Wallet Management Menu")
		walletManagementMenu.Add(getGetBalanceMenuItem(service))
		walletManagementMenu.Add(getDisplayCardsMenuItem(service))
		walletManagementMenu.Add(getAddCardMenuItem(service))
		walletManagementMenu.Add(getDepositMenuItem(service))
		walletManagementMenu.Add(getWithDrawMenu(service))
	}
	return walletManagementMenu
}

func getManageAddressesMenu(service services.Service) MenuComponent {
	// if manageAddresses == nil {

	// }
	return manageAddresses
}

func getEditInfoMenu(service services.Service) MenuComponent {
	// if editInfoMenu == nil {
	// 	//editInfoMenu = New
	// }
	return editInfoMenu
}

func getProfileManagementMenu(service services.Service) MenuComponent {
	if profileManagementMenu == nil {
		profileManagementMenu = NewMenu("Profile Management Menu")
		profileManagementMenu.Add(getEditInfoMenu(service))
		profileManagementMenu.Add(getManageAddressesMenu(service))
	}
	return profileManagementMenu
}

func getFoodOrderMenu(service services.Service) MenuComponent {
	if foodOrderMenu == nil {
		foodOrderMenu = NewMenu("Food Order Menu")
	}
	return foodOrderMenu
}

func getRestaurantManagementMenu(service services.Service) MenuComponent {
	if restaurantManagementMenu == nil {
		restaurantManagementMenu = NewMenu("Restaurant Management Menu")
	}
	return restaurantManagementMenu
}

func getHomeMenu(service services.Service) MenuComponent {
	if homeMenu == nil {
		homeMenu = NewMenu("Home Menu")
		homeMenu.Add(getProfileManagementMenu(service))
		homeMenu.Add(getWalletManagementMenu(service))
		homeMenu.Add(getFoodOrderMenu(service))
		homeMenu.Add(getRestaurantManagementMenu(service))
		homeMenu.Add(getLogoutMenuItem(service))

	}
	return homeMenu
}

func getLogoutMenuItem(service services.Service) MenuComponent {
	if logoutMenuItem == nil {
		logoutMenuItem = NewLogoutMenuItem("Logout", commands.NewLogoutCommand(service), GetMainMenu(service))
	}
	return logoutMenuItem
}

func getLoginMenuItem(service services.Service) MenuComponent {
	if loginMenuItem == nil {
		loginMenuItem = NewLoginMenuItem("Login Menu", commands.NewLoginCommand(service), getHomeMenu(service))
	}
	return loginMenuItem
}

func getRegisterMenuItem(service services.Service) MenuComponent {
	if registerMenuItem == nil {
		registerMenuItem = NewRegisterMenuItem("Register Menu", commands.NewRegisterCommand(service), getHomeMenu(service))
	}
	return registerMenuItem
}

func GetMainMenu(service services.Service) MenuComponent {
	if mainMenu == nil {
		mainMenu = NewMenu("Main Menu")
		mainMenu.Add(getLoginMenuItem(service))
		mainMenu.Add(getRegisterMenuItem(service))
	}
	return mainMenu
}
