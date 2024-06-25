package mappers

import (
	"server/internal/models/restaurant/restaurant"
	"server/pkg/adapters/storage/entities"
)

func RestaurantEntityToDomain(entity *entities.Restaurant) *restaurant.Restaurant {
	return &restaurant.Restaurant{
		ID:    entity.ID,
		Name:  entity.Name,
		Phone: entity.Phone,
	}
}

func RestaurantDomainToEntity(domainRestaurant *restaurant.Restaurant) *entities.Restaurant {
	return &entities.Restaurant{
		Name:  domainRestaurant.Name,
		Phone: domainRestaurant.Phone,
	}
}
