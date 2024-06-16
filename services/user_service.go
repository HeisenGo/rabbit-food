package services

import (
	"rabbit-food/internal/models/user"
	"rabbit-food/pkg/adapters/storage/entities"
)

type UserService struct {
	userOps *user.Ops
}

func NewUserService(userOps *user.Ops) *UserService {
	return &UserService{
		userOps: userOps,
	}
}

// func (s *UserService) CreateUser(ctx context.Context, user *user.User) error {
func (s *UserService) CreateUser(user *user.User) (*entities.User, error) {
	return s.userOps.Create(user)
}
