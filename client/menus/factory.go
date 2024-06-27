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
	displayBallanceMenu    MenuComponent
	diaplayCardsMenu       MenuComponent
	connectACreditCardMenu MenuComponent
	depositMenuItem        MenuComponent
	withDrawMenuItem       MenuComponent
)

func getDisplayCardsMenu(service services.Service) MenuComponent {
	if diaplayCardsMenu == nil {
		diaplayCardsMenu = NewDisplayCardsMenuItem("Cards", commands.NewDisplayCardsCommand(service), getWalletManagementMenu(service))
	}
	return diaplayCardsMenu
}

func getDisplayBalanceMenu(service services.Service) MenuComponent {
	if displayBallanceMenu == nil {
		//displayBallanceMenu = NewDisplayBalanceMenuItem("Display Ballance", commands.NewDisplayBallanceCommand(service), getWalletManagementMenu(service))
	}
	return displayBallanceMenu
}

func getAddCardMenuItem(service services.Service) MenuComponent {
	if connectACreditCardMenu == nil {
		connectACreditCardMenu = NewAddCardMenuItem("Connect A New Credit Card", commands.NewAddCardCommand(service), getWalletManagementMenu(service))
	}
	return connectACreditCardMenu
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
		//walletManagementMenu.Add(getDisplayBalanceMenu(service))
		walletManagementMenu.Add(getDisplayCardsMenu(service))
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

func getEditInfoMenue(service services.Service) MenuComponent {
	return NewMenu("Edit Info Menu")
}

func getProfileManagementMenu(service services.Service) MenuComponent {
	if profileManagementMenu == nil {
		profileManagementMenu = NewMenu("Profile Management Menu")
		profileManagementMenu.Add(getEditInfoMenue(service))
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
		homeMenu.Add(getLogoutMenuItem(service)) // TODO: remove token from context & add context to flow

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

// ///////////////////////////////////////////////////////////////////////
//func gethomeMenu(service services.Service) MenuComponent {
//	return &Menu{
//		Name: "Profile Management Menu",
//		MenuItems: []MenuComponent{
//			getWalletManagementMenu(service),
//		},
//	}
//}
//
//func getRegisterMenuItem(service services.Service) MenuComponent {
//	return NewRegisterMenuItem("Register", commands.NewRegisterCommand(service), gethomeMenu(service))
//}

// Similarly, get functions for other menus
