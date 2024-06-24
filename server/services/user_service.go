package services

import (
	"context"
	"server/internal/models/user"
)

type UserService struct {
	userOps *user.Ops
}

func NewUserService(userOps *user.Ops) *UserService {
	return &UserService{
		userOps: userOps,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *user.User) (*user.User, error) {
	return s.userOps.Create(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*user.User, error) {
	return s.userOps.GetUserByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx context.Context, user *user.User) (*user.User, error) {
	return s.userOps.Update(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	return s.userOps.Delete(ctx, id)
}
