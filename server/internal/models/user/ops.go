package user

import (
	"context"
	"gorm.io/gorm"
	"server/pkg/adapters/storage/entities"
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
	return o.repo.Create(user)
}
