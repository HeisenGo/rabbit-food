package user

import (
	"context"
	"server/pkg/adapters/storage/entities"
	"server/pkg/utils/users"
dberr"server/internal/errors/users"
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
	hashedPass, err := users.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.SetPassword(hashedPass)
	return o.repo.Create(ctx,user)
}
func (o *Ops) GetUserByEmailAndPassword(ctx context.Context, email, password string) (*User, error) {
	user, err := o.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, dberr.ErrUserNotFound
	}

	if !user.PasswordIsValid(password) {
		return nil, dberr.ErrInvalidPassword
	}

	return user, nil
}
func (o *Ops) GetUserByPhoneAndPassword(ctx context.Context, phone, password string) (*User, error) {
	user, err := o.repo.GetByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, dberr.ErrUserNotFound
	}

	if !user.PasswordIsValid(password) {
		return nil, dberr.ErrInvalidPassword
	}

	return user, nil
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
