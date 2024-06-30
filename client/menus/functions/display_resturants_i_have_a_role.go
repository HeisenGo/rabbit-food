package functions

import (
	"client/models"
	"fmt"
	"strings"
)

func DisplayRestaurantsWithAddress(restaurants []*models.Restaurant) {
	for i, rest := range restaurants {
		strSlice := []string{rest.Name, rest.Phone, rest.Address.City, rest.Address.AddressLine}
		str := strings.Join(strSlice, " ")
		fmt.Printf("\n\t%d. %s", i+1, str)
		fmt.Println("\n\t--------------------------------------------------------------")
	}
}

func DisplayMenus(menus []*models.RestaurantMenu) {
	for i, m := range menus {
		fmt.Printf("\n\t%d. %s", i+1, m.Name)
		fmt.Println("\n\t--------------------------------------------------------------")
	}
}
