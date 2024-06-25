package mappers

import (
	"server/internal/models/restaurant/motor"
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

func MotorDomainToEntity(domianMotor *motor.Motor) *entities.Motor {
	return &entities.Motor{
		Name:  domianMotor.Name,
		Speed: domianMotor.Speed,
	}
}

func MotorEntityToDomain(entiryMotor *entities.Motor) *motor.Motor {
	return &motor.Motor{
		Name:         entiryMotor.Name,
		Speed:        entiryMotor.Speed,
		RestaurantID: entiryMotor.RestaurantID,
	}
}
