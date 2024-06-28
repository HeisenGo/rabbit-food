package mappers

import (
	"fmt"
	"server/internal/models/restaurant/restaurant"
	"server/pkg/adapters/storage/entities"
)

func RestaurantEntityToDomain(entity *entities.Restaurant) *restaurant.Restaurant {
  	domainAddress := RestaurantAddressEntityToDomain(entity.Address)
	return &restaurant.Restaurant{
		ID:         entity.ID,
		Name:       entity.Name,
		Phone:      entity.Phone,
    Address:    domainAddress,
		Categories: BatchRestaurantCategoryEntityToDomain(entity.Categories),
}

func RestaurantDomainToEntity(domainRestaurant *restaurant.Restaurant) *entities.Restaurant {
	entAddress := AddressDomainToEntity(domainRestaurant.Address)
	return &entities.Restaurant{
		Name:    domainRestaurant.Name,
		Phone:   domainRestaurant.Phone,
		Address: entAddress,
	}
}

func BatchRestaurantCategoryEntityToDomain(entities []*entities.RestaurantCategory) []*restaurant.RestaurantCategory {
	var domainCategories []*restaurant.RestaurantCategory
	for _, e := range entities {
		domainCategories = append(domainCategories, &restaurant.RestaurantCategory{ID: e.ID, Name: e.Name})
	}
	return domainCategories
}
