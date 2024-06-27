package mappers
import (
	"server/internal/models/address"
	"server/pkg/adapters/storage/entities"
)

func AddressEntityToDomain(entity *entities.Address) *address.Address {
	return &address.Address{			
		Addressline : entity.Addressline,		
		Cordinates  : entity.Cordinates ,			
		Types  		: entity.Types		,	
		City 		: entity.City 		,	
	}
}

func AddressDomainToEntity(domainAddress *address.Address) *entities.Address {
	return &entities.Address{
		Addressline:domainAddress.Addressline,
		Cordinates: domainAddress.Cordinates,
		Types:  	domainAddress.Types,
		City:  		domainAddress.City,
	}
}


