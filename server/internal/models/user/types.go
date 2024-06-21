package user

import (
	"context"
	"time"
)

type Repo interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByPhone(ctx context.Context, phone string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type User struct {
	ID        uint
	Phone     string
	Email     *string
	FirstName string
	LastName  string
	Password  string
	BirthDate time.Time
	IsAdmin   bool
}

func NewUser(phone string, email *string, password string) *User {
	return &User{
		Phone:    phone,
		Email:    email,
		Password: password,
	}
}

func (u *User) SetEmail(email string) {
	u.Email = &email
}
func (u *User) SetPhone(phone string) {
	u.Phone = phone
}
func (u *User) SetPassword(password string) {
	u.Password = password
}
