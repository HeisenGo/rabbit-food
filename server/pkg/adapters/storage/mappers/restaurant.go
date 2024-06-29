package mappers

import (
	"fmt"
	"server/internal/models/restaurant/restaurant"
	"server/pkg/adapters/storage/entities"
)

func RestaurantEntityToDomain(entity *entities.Restaurant) *restaurant.Restaurant {
	domainAddress := RestaurantAddressEntityToDomain(entity.Address)
	r := &restaurant.Restaurant{
		ID:      entity.ID,
		Name:    entity.Name,
		Phone:   entity.Phone,
		Address: domainAddress,
	}
	fmt.Print(r)
	return r
}

func RestaurantDomainToEntity(domainRestaurant *restaurant.Restaurant) *entities.Restaurant {
	entAddress := AddressDomainToEntity(domainRestaurant.Address)
	return &entities.Restaurant{
		Name:    domainRestaurant.Name,
		Phone:   domainRestaurant.Phone,
		Address: entAddress,
	}
}
