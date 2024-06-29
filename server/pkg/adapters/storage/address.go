package storage

import (
	"context"
	"gorm.io/gorm"
	"server/internal/models/address"
	"server/pkg/adapters/storage/mappers"
	"server/pkg/utils"
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
	newAddress := mappers.AddressDomainToEntity(address)
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	newAddress.UserID = &userID
	err = r.db.Create(&newAddress).Error
	if err != nil {
		return nil, err
	}
	// Convert entity model back to domain model
	return mappers.AddressEntityToDomain(newAddress), nil
}
