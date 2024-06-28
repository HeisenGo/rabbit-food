package mappers

import (
	"server/internal/models/address"
	"server/pkg/adapters/storage/entities"
)

func AddressEntityToDomain(entity *entities.Address) *address.Address {
	return &address.Address{
		UserID:      entity.UserID,
		AddressLine: entity.AddressLine,
		Coordinates: entity.Coordinates,
		Types:       entity.Types,
		City:        entity.City,
	}
}

func AddressDomainToEntity(domainAddress *address.Address) *entities.Address {
	return &entities.Address{
		AddressLine: domainAddress.AddressLine,
		Coordinates: domainAddress.Coordinates,
		Types:       domainAddress.Types,
		City:        domainAddress.City,
	}
}
