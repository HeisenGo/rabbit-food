package menus

import (
	"client/commands"
	"client/services"
)

var (
	mainMenu                 *Menu
	registerMenuItem         MenuComponent
	loginMenuItem            MenuComponent
	homeMenu                 *Menu // TODO: bayad be interface vabaste bashe ye fekri bokon
	profileManagementMenu    MenuComponent
	walletManagementMenu     MenuComponent
	foodOrderMenu            MenuComponent
	restaurantManagementMenu MenuComponent
	logoutMenuItem           MenuComponent
)

func getWalletManagementMenu(service services.Service) MenuComponent {
	if walletManagementMenu == nil {
		walletManagementMenu = NewWalletMenuItem("Wallet Management Menu", commands.NewGetWalletCommand(service), nil)
	}
	return walletManagementMenu
}

func getProfileManagementMenu(service services.Service) MenuComponent {
	if profileManagementMenu == nil {
		profileManagementMenu = NewMenu("Profile Management Menu")
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
