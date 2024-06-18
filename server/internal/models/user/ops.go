package user

import (
	"context"
	"server/pkg/adapters/storage/entities"
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
	// validation
	validateUserRegistration(user)
	return o.repo.Create(user)
}

func validateUserRegistration(user *User) error {
	err := validations.ValidateEmail(user.Email)
	if err != nil {
		return err
	}
	if err = validations.ValidatePasswordWithFeedback(user.Password); err != nil {
		return err
	}
	if err = validations.ValidatePoneNumber(user.Phone); err != nil {
		return err
	}
	return nil
}
