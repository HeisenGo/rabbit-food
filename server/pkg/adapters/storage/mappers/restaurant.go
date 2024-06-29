package mappers

import (
	"fmt"
	"server"
	"server/internal/models/address"
	"server/internal/models/restaurant/motor"
	"server/internal/models/restaurant/restaurant"
	"server/pkg/adapters/storage/entities"
)

func RestaurantEntityToDomain(entity *entities.Restaurant) *restaurant.Restaurant {
	fmt.Println(entity)
	var domainAddress *address.Address
	if entity.Address == nil {
		domainAddress = address.NewAddress("", address.Coordinates{}, server.RestaurantAddressType, "")
	} else {
		domainAddress = RestaurantAddressEntityToDomain(entity.Address)
	}
	r := &restaurant.Restaurant{
		ID:      entity.ID,
		Name:    entity.Name,
		Phone:   entity.Phone,
		Address: domainAddress,
	}
	fmt.Print(r)
	return r
}

func RestaurantEntityAddressNameLineToDomain(entity *entities.Restaurant) *restaurant.Restaurant {
	var domainAddress *address.Address
	if entity.Address == nil {
		domainAddress = address.NewAddress("", address.Coordinates{}, server.RestaurantAddressType, "")
	} else {
		domainAddress = RestaurantAddressNameLineEntityToDomain(entity.Address)
	}
	r := &restaurant.Restaurant{
		ID:      entity.ID,
		Name:    entity.Name,
		Phone:   entity.Phone,
		Address: domainAddress,
	}
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

func MotorDomainToEntity(domainMotor *motor.Motor) *entities.Motor {
	return &entities.Motor{
		Name:  domainMotor.Name,
		Speed: domainMotor.Speed,
	}
}

func MotorEntityToDomain(entityMotor *entities.Motor) *motor.Motor {
	return &motor.Motor{
		Name:         entityMotor.Name,
		Speed:        entityMotor.Speed,
		RestaurantID: entityMotor.RestaurantID,
	}
}
