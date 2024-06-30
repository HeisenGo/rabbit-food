package storage

import (
	"context"
	"errors"
	"strings"

	"server/internal/errors/users"
	"server/internal/models/user"
	"server/pkg/adapters/storage/entities"
	"server/pkg/adapters/storage/mappers"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func (r *userRepo) Create(ctx context.Context, user *user.User) (*user.User, error) {
	newUser := mappers.UserDomainToEntity(user)
	err := r.db.Create(&newUser).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, users.ErrUserExists
		}
		return nil, err
	}
	createdUser := mappers.UserEntityToDomain(newUser)
	return createdUser, nil
}

func (r *userRepo) GetByPhone(ctx context.Context, phone string) (*user.User, error) {
	var userEntity entities.User
	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("phone = ?", phone).First(&userEntity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, users.ErrUserNotFound
		}
		return nil, err
	}
	return mappers.UserEntityToDomain(&userEntity), nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var userEntity entities.User
	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("email = ?", email).First(&userEntity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, users.ErrUserNotFound
		}
		return nil, err
	}
	return mappers.UserEntityToDomain(&userEntity), nil
}

func (r *userRepo) GetByID(ctx context.Context, id uint) (*user.User, error) {
	var userEntity entities.User
	err := r.db.WithContext(ctx).Model(&entities.User{}).Where("id = ?", id).First(&userEntity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, users.ErrUserNotFound
		}
		return nil, err
	}
	return mappers.UserEntityToDomain(&userEntity), nil
}

func (r *userRepo) Update(ctx context.Context, user *user.User) (*user.User, error) {
	existingUser := mappers.UserDomainToEntity(user)
	err := r.db.Save(&existingUser).Error
	if err != nil {
		return nil, err
	}
	updatedUser := mappers.UserEntityToDomain(existingUser)
	return updatedUser, nil
}

func NewUserRepo(db *gorm.DB) user.Repo {
	return &userRepo{
		db: db,
	}
}
