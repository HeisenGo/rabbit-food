package services

import (
	"context"
	"server/internal/models/user"
	"server/pkg/adapters/storage/entities"
)

type UserService struct {
	userOps *user.Ops
}

func NewUserService(userOps *user.Ops) *UserService {
	return &UserService{
		userOps: userOps,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *user.User) (*entities.User, error) {
	return s.userOps.Create(ctx, user)
}
func (s *UserService)LoginUserByEmail(ctx context.Context, email,pass string)(*user.User,error){
	return s.userOps.GetUserByEmailAndPassword(ctx,email,pass)
}
func (s *UserService)LoginUserByPhone(ctx context.Context, phone,pass string)(*user.User,error){
	return s.userOps.GetUserByPhoneAndPassword(ctx,phone,pass)
}