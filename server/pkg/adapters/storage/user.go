package storage

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"server/internal/models/user"
	"server/pkg/adapters/storage/entities"
	"server/pkg/adapters/storage/mappers"
)

type userRepo struct {
	db *gorm.DB
}

func (r *userRepo) Create(ctx context.Context, user *user.User) (*user.User, error) {
	newUser := mappers.UserDomainToEntity(user)
	err := r.db.Create(&newUser).Error
	if err != nil {
		return nil, err
	}
	createdUser := mappers.UserEntityToDomain(newUser)
	return createdUser, nil
}

func (r *userRepo) GetByPhone(ctx context.Context, phone string) (*user.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("Phone = ?", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mappers.UserEntityToDomain(&user), nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var user entities.User
	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("Email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mappers.UserEntityToDomain(&user), nil
}

func NewUserRepo(db *gorm.DB) user.Repo {
	return &userRepo{
		db: db,
	}
}
