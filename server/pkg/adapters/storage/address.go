package storage

import (
	"context"
	"server/internal/models/address"
	"server/pkg/adapters/storage/entities"
	"server/pkg/adapters/storage/mappers"
	"server/pkg/utils"
	"gorm.io/gorm"
)

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepo(db *gorm.DB) address.Repo {
	return &addressRepo{
		db: db,
	}
}
func (r *addressRepo) Create(ctx context.Context, address *address.Address) (*address.Address, error) {
	// Convert domain model to entity model
	newAddress := mappers.AddressDomainToEntity(address)
	userid, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	newAddress.UserID = userid
	// Extract cordinates
	longitude := address.Cordinates[0]
	latitude := address.Cordinates[1]
    var result entities.Address
    resulted := r.db.Raw(`
        INSERT INTO addresses (addressline, cordinates, types, city, user_id, created_at, updated_at)
        VALUES ($1, ST_SetSRID(ST_MakePoint($2, $3), 4326), $4, $5, $6, NOW(), NOW())
        RETURNING id, addressline, ST_X(cordinates::geometry) AS longitude, ST_Y(cordinates::geometry) AS latitude, types, city, user_id, created_at, updated_at`,
        newAddress.Addressline,longitude, latitude, newAddress.Types, newAddress.City, newAddress.UserID).
        Scan(&result)
    if resulted.Error != nil {
        return nil, resulted.Error
    }
	// Convert entity model back to domain model
	return mappers.AddressEntityToDomain(&result), nil
}