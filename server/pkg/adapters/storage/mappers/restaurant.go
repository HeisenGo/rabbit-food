package mappers

import (
	"server/internal/models/restaurant/restaurant"
	"server/pkg/adapters/storage/entities"
)

func RestaurantEntityToDomain(entity *entities.Restaurant) *restaurant.Restaurant {
	domainAddress := AddressEntityToDomain(entity.Address)
	return &restaurant.Restaurant{
		ID:      entity.ID,
		Name:    entity.Name,
		Phone:   entity.Phone,
		Address: domainAddress,
	}
}

func RestaurantDomainToEntity(domainRestaurant *restaurant.Restaurant) *entities.Restaurant {
	entAddress := AddressDomainToEntity(domainRestaurant.Address)
	return &entities.Restaurant{
		Name:    domainRestaurant.Name,
		Phone:   domainRestaurant.Phone,
		Address: entAddress,
	}
}
