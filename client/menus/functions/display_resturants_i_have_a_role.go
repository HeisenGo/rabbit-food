package functions

import (
	"client/models"
	"fmt"
	"strings"
)

func DisplayRestaurantsWithAddress(restaurants []*models.Restaurant) {
	for _, rest := range restaurants {
		strSlice := []string{rest.Name, rest.Phone, rest.Address.City, rest.Address.AddressLine}
		str := strings.Join(strSlice, " ")
		fmt.Println(str)
	}
}
