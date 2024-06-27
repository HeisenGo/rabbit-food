package user

import (
	"context"
	userErrors "server/internal/errors/users"
	"server/pkg/utils/users"

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

func (o *Ops) Create(ctx context.Context, user *User) (*User, error) {
	err := validateUserRegistration(user)
	if err != nil {
		return nil, err
	}
	hashedPass, err := users.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.SetPassword(hashedPass)
	return o.repo.Create(ctx, user)
}
func (o *Ops) GetUser(ctx context.Context, phoneOrEmail, password string) (*User, error) {
	var user *User
	if users.ValidatePhone(phoneOrEmail) != nil {
		if users.ValidateEmail(phoneOrEmail) != nil {
			return nil, userErrors.ErrUserPassDoesNotMatch
		}
		email := phoneOrEmail
		user, _ = o.repo.GetByEmail(ctx, email)
	} else {
		phone := phoneOrEmail
		user, _ = o.repo.GetByPhone(ctx, phone)
	}

	if user == nil {
		return user, userErrors.ErrUserPassDoesNotMatch
	}

	if err := users.CheckPasswordHash(password, user.Password); err != nil {
		return nil, userErrors.ErrUserPassDoesNotMatch
	}
	return user, nil
}

func validateUserRegistration(user *User) error {
	if err := users.ValidatePhone(user.Phone); err != nil {
		return err
	}
	if user.Email != nil {
		err := users.ValidateEmail(*user.Email)
		if err != nil {
			return err
		}
	}

	if err := users.ValidatePasswordWithFeedback(user.Password); err != nil {
		return err
	}
	return nil
}

func (o *Ops) GetByID(ctx context.Context, userID uint) (*User, error) {
	var user User
	if err := o.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (o *Ops) Update(ctx context.Context, user *User) (*User, error) {
	if err := o.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
