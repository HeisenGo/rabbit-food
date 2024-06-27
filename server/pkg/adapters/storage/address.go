package storage

import (
	"context"
	"errors"
	"server/internal/errors/users"
	"server/internal/models/address"
	"server/pkg/adapters/storage/entities"
	"server/pkg/adapters/storage/mappers"
	//"server/pkg/utils"
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
func (r *addressRepo) Create(ctx context.Context, address *address.Address, userID uint) (*address.Address, error) {
	// Convert domain model to entity model
	newAddress := mappers.AddressDomainToEntity(address)
	newAddress.UserID = &userID

	// Extract cordinates
	longitude := address.Cordinates[0]
	latitude := address.Cordinates[1]

	// Use raw SQL to insert the address with cordinates as geography type
	err := r.db.Exec(`
		INSERT INTO addresses (addressline, cordinates, types, city, user_id, created_at, updated_at)
		VALUES (?, ST_SetSRID(ST_MakePoint(?, ?), 4326), ?, ?, ?, NOW(), NOW())`,
		newAddress.Addressline, longitude, latitude, newAddress.Types, newAddress.City, userID).Error

	if err != nil {
		return nil, err
	}

	// Retrieve the created address to return
	var createdAddress *entities.Address
	err = r.db.Raw(`
		SELECT id, addressline, ST_X(cordinates::geometry) AS longitude, ST_Y(cordinates::geometry) AS latitude, types, city, user_id, created_at, updated_at
		FROM addresses
		WHERE addressline = ? AND user_id = ?
		LIMIT 1`, newAddress.Addressline, userID).Scan(&createdAddress).Error
	if err != nil {
		return nil, err
	}

	// Convert entity model back to domain model
	return mappers.AddressEntityToDomain(createdAddress), nil
}
func (r *addressRepo) GetByUser(ctx context.Context, userID uint ) (*address.Address, error) {
	var addressEntity entities.Address
	err := r.db.WithContext(ctx).Model(&entities.Address{}).Where("userid = ?", userID).First(&addressEntity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, users.ErrAddressNotFound
		}
		return nil, err
	}
	return mappers.AddressEntityToDomain(&addressEntity), nil
}

func (r *addressRepo) GetByRestaurant(ctx context.Context, name string) (*address.Address, error) {
	var addressEntity entities.Address
	err := r.db.WithContext(ctx).Model(&entities.Address{}).Where("name = ?", name).First(&addressEntity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//Change to restaurant error
			return nil, users.ErrAddressNotFound
		}
		return nil, err
	}
	return mappers.AddressEntityToDomain(&addressEntity), nil
}