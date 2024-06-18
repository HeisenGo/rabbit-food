package user

import (
	"context"
	"gorm.io/gorm"
	"server/pkg/adapters/storage/entities"
	"server/pkg/utils/users"
)

type Ops struct {
	db   *gorm.DB
	repo Repo
}

func NewUserOps(db *gorm.DB, repo Repo) *Ops {
	return &Ops{
		db:   db,
		repo: repo,
	}
}

func (o *Ops) Create(ctx context.Context, user *User) (*entities.User, error) {
	err := validateUserRegistration(user)
	if err != nil {
		return nil, err
	}
	hashedPass, err := users.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.SetPassword(hashedPass)
	return o.repo.Create(user)
}

func validateUserRegistration(user *User) error {
	if err := users.ValidatePhoneNumber(user.Phone); err != nil {
		return err
	}

	if user.Email != "" {
		err := users.ValidateEmail(user.Email)
		if err != nil {
			return err
		}
	}

	if err := users.ValidatePasswordWithFeedback(user.Password); err != nil {
		return err
	}
	return nil
}
