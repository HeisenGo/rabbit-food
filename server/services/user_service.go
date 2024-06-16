package services

import (
	user2 "rabbit-food/server/internal/models/user"
	"rabbit-food/server/pkg/adapters/storage/entities"
)

type UserService struct {
	userOps *user2.Ops
}

func NewUserService(userOps *user2.Ops) *UserService {
	return &UserService{
		userOps: userOps,
	}
}

// func (s *UserService) CreateUser(ctx context.Context, user *user.User) error {
func (s *UserService) CreateUser(user *user2.User) (*entities.User, error) {
	return s.userOps.Create(user)
}
