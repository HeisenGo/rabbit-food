package user

import (
	"context"
	"gorm.io/gorm"
	"server/pkg/adapters/storage/entities"
	"server/pkg/utils/validations"
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
// func (o *Ops) Create(ctx context.Context, user *User) error {
func (o *Ops) Create(user *User) (*entities.User, error) {
	err := validations.ValidateUserRegistration(user)
	if err != nil {
		return nil, err
	}
	return o.repo.Create(user)
}
