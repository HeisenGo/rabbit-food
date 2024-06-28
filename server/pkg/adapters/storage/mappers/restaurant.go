package mappers

import (
	"server/internal/models/restaurant/restaurant"
	"server/pkg/adapters/storage/entities"
)

func RestaurantEntityToDomain(entity *entities.Restaurant) *restaurant.Restaurant {
	return &restaurant.Restaurant{
		ID:         entity.ID,
		Name:       entity.Name,
		Phone:      entity.Phone,
		Categories: BatchRestaurantCategoryEntityToDomain(entity.Categories),
	}
}

func RestaurantDomainToEntity(domainRestaurant *restaurant.Restaurant) *entities.Restaurant {
	return &entities.Restaurant{
		Name:  domainRestaurant.Name,
		Phone: domainRestaurant.Phone,
	}
}

func BatchRestaurantCategoryEntityToDomain(entities []*entities.RestaurantCategory) []*restaurant.RestaurantCategory {
	var domainCategories []*restaurant.RestaurantCategory
	for _, e := range entities {
		domainCategories = append(domainCategories, &restaurant.RestaurantCategory{ID: e.ID, Name: e.Name})
	}
	return domainCategories
}
