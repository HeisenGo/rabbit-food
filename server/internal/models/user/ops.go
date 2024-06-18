package user

import (
	"context"
	"server/pkg/adapters/storage/entities"
	"server/pkg/utils/hash"
	"server/pkg/utils/validations"

	"gorm.io/gorm"
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
	hashedPass, err := hash.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.SetPassword(hashedPass)
	return o.repo.Create(user)
}

func validateUserRegistration(user *User) error {
	if err := validations.ValidatePhoneNumber(user.Phone); err != nil {
		return err
	}

	if user.Email != "" {
		err := validations.ValidateEmail(user.Email)
		if err != nil {
			return err
		}
	}

	if err := validations.ValidatePasswordWithFeedback(user.Password); err != nil {
		return err
	}
	return nil
}
