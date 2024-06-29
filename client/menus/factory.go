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
	restaurantManagementMenu *Menu
	logoutMenuItem           MenuComponent
	//////       Profile ManagementMeu /////
	///// Info to display and last updated address to display
	editInfoMenu    *Menu
	manageAddresses *Menu
	editAddress     MenuComponent
	deleteAddress   MenuComponent
	addNewAddress   MenuComponent

	/// wallet management menu ////
	displayBallanceMenuItem MenuComponent
	displayCardsMenuItem    MenuComponent
	AddCreditCardMenuItem   MenuComponent
	depositMenuItem         MenuComponent
	withDrawMenuItem        MenuComponent
	// restaurant management
	createRestaurantMenuItem         MenuComponent
	showMyFunctionalRestaurants      MenuComponent
	addCategoryToRestaurantsMenuItem MenuComponent
	addMenuToRestaurantMenuItem      MenuComponent
	addFoodToRestaurantMenuItem      MenuComponent
)

func getAddMenuToRestaurantMenuItem(service services.Service) MenuComponent {
	if addMenuToRestaurantMenuItem == nil {
		addMenuToRestaurantMenuItem = NewAddMenuToRestaurantMenuItem("Add Menu", commands.NewAddMenuToRestaurantCommand(service), commands.NewGetRestaurantsIHaveARoleCommand(service), commands.NewGetCategoriesOfRestaurantCommand(service), getRestaurantManagementMenu(service))
	}
	return addMenuToRestaurantMenuItem
}

func getShowMyFunctionalRestaurants(service services.Service) MenuComponent {
	if showMyFunctionalRestaurants == nil {
		showMyFunctionalRestaurants = NewGetRestaurantsIHaveARoleInitMenuItem("Restaurant With a role", commands.NewGetRestaurantsIHaveARoleCommand(service), getRestaurantManagementMenu(service))
	}
	return showMyFunctionalRestaurants
}

func getAddCategoryToRestaurantsMenuItem(service services.Service) MenuComponent {
	if addCategoryToRestaurantsMenuItem == nil {
		addCategoryToRestaurantsMenuItem = NewAddCategoryToRestaurantMenuItem("Add Category to Restaurant", commands.NewAddCategoryToRestaurantCommand(service), commands.NewGetRestaurantsIHaveARoleCommand(service), getRestaurantManagementMenu(service))
	}
	return addCategoryToRestaurantsMenuItem
}

func getCreateRestaurantMenuItem(service services.Service) MenuComponent {
	if createRestaurantMenuItem == nil {
		createRestaurantMenuItem = NewCreateRestaurantMenuItem("Create Restaurant", commands.NewCreateRestaurantCommand(service), getRestaurantManagementMenu(service))
	}
	return createRestaurantMenuItem
}

func getDisplayCardsMenuItem(service services.Service) MenuComponent {
	if displayCardsMenuItem == nil {
		displayCardsMenuItem = NewDisplayCardsMenuItem("Cards", commands.NewDisplayCardsCommand(service), getWalletManagementMenu(service))
	}
	return displayCardsMenuItem
}

func getDisplayBalanceMenuItem(service services.Service) MenuComponent {
	if displayBallanceMenuItem == nil {
		displayBallanceMenuItem = NewGetWalletMenuItem("Display Ballance", commands.NewGetWalletCommand(service), getWalletManagementMenu(service))
	}
	return displayBallanceMenuItem
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
		walletManagementMenu.Add(getDisplayBalanceMenuItem(service))
		walletManagementMenu.Add(getDisplayCardsMenuItem(service))
		walletManagementMenu.Add(getAddCardMenuItem(service))
		walletManagementMenu.Add(getDepositMenuItem(service))
		walletManagementMenu.Add(getWithDrawMenu(service))
		walletManagementMenu.Add(getLogoutMenuItem(service))

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
		profileManagementMenu.Add(getLogoutMenuItem(service))
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
		restaurantManagementMenu.Add(getCreateRestaurantMenuItem(service))
		restaurantManagementMenu.Add(getAddCategoryToRestaurantsMenuItem(service))
		restaurantManagementMenu.Add(getAddMenuToRestaurantMenuItem(service))
		restaurantManagementMenu.Add(getShowMyFunctionalRestaurants(service))
		restaurantManagementMenu.Add(getLogoutMenuItem(service))
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
