package user

import (
	"server/pkg/adapters/storage/entities"
	"time"
)

type Repo interface {
	Create(user *User) (*entities.User, error)
	//Create(ctx context.Context, user *User) error
	//GetByID(ctx context.Context, id uint) (*User, error)
	//GetByEmail(ctx context.Context, email string) (*User, error)
}

type User struct {
	ID        uint
	Phone     string
	Email     string
	FirstName string
	LastName  string
	Password  string
	BirthDate time.Time
}

func NewUser(phone, email, password string) *User {
	return &User{
		Phone:    phone,
		Email:    email,
		Password: password,
	}
}

func (u *User) SetPassword(password string) {
	u.Password = password
}
