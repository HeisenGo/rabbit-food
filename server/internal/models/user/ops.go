package user

import (
	"context"
	"errors"
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

func (o *Ops) GetOperatorUser(ctx context.Context, phoneOrEmail string) (*User, error) {
	var user *User
	if users.ValidatePhone(phoneOrEmail) != nil {
		if users.ValidateEmail(phoneOrEmail) != nil {
			return nil, userErrors.ErrUserNotFound
		}
		email := phoneOrEmail
		user, _ = o.repo.GetByEmail(ctx, email)
	} else {
		phone := phoneOrEmail
		user, _ = o.repo.GetByPhone(ctx, phone)
	}

	if user == nil {
		return user, userErrors.ErrUserNotFound
	}
	return user, nil
}

func (o *Ops) GetByID(ctx context.Context, id uint) (*User, error) {
	return o.repo.GetByID(ctx, id)
}

func (o *Ops) UpdateField(ctx context.Context, id uint, fieldName string, value interface{}) (*User, error) {
	user, err := o.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	switch fieldName {
	case "first_name":
		user.FirstName = value.(string)
	case "last_name":
		user.LastName = value.(string)
	case "email":
		email := value.(string)
		user.Email = &email
	case "phone":
		user.Phone = value.(string)
	case "password":
		user.Password = value.(string)
	default:
		return nil, errors.New("invalid field name")
	}

	return o.repo.Update(ctx, user)
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
