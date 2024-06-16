package storage

import (
	"gorm.io/gorm"
	"server/internal/models/user"
	"server/pkg/adapters/storage/entities"
)

type userRepo struct {
	db *gorm.DB
}

func (r *userRepo) Create(user *user.User) (*entities.User, error) {
	newUser := &entities.User{}
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
