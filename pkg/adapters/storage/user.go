package storage

import (
	"gorm.io/gorm"
	"rabbit-food/internal/models/user"
	"rabbit-food/pkg/adapters/storage/entities"
)

type userRepo struct {
	db *gorm.DB
}

func (r *userRepo) Create(user *user.User) (*entities.User, error) {
	var newUser *entities.User
	newUser.Phone = user.Phone
	newUser.Email = user.Email
	newUser.Password = user.Password
	err := r.db.Create(&newUser).Error
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func NewUserRepo(db *gorm.DB) user.Repo {
	return &userRepo{
		db: db,
	}
}
