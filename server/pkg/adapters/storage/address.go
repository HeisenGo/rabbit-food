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
	newAddress := mappers.AddressDomainToEntity(address)
	newAddress.UserID = &userID
	err := r.db.Create(&newAddress).Error
	if err != nil {
		return nil, err
	}
	createdAddress := mappers.AddressEntityToDomain(newAddress)
	return createdAddress, nil
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