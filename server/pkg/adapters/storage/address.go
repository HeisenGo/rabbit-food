package storage

import (
	"context"
	"errors"

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
	newAddress := mappers.AddressDomainToEntity(address)
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	newAddress.UserID = &userID
	err = r.db.Create(newAddress).Error
	if err != nil {
		return nil, err
	}
	return mappers.AddressEntityToDomain(newAddress), nil
}

func (r *addressRepo) GetByID(ctx context.Context, id uint) (*address.Address, error) {
	var addressEntity entities.Address
	err := r.db.WithContext(ctx).First(&addressEntity, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("address not found")
		}
		return nil, err
	}
	return mappers.AddressEntityToDomain(&addressEntity), nil
}

func (r *addressRepo) Update(ctx context.Context, address *address.Address) (*address.Address, error) {
	existingAddress := mappers.AddressDomainToEntity(address)
	err := r.db.Save(existingAddress).Error
	if err != nil {
		return nil, err
	}
	return mappers.AddressEntityToDomain(existingAddress), nil
}

func (r *addressRepo) Delete(ctx context.Context, id uint) error {
	err := r.db.WithContext(ctx).Delete(&entities.Address{}, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("address not found")
		}
		return err
	}
	return nil
}
