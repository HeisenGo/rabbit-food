package services

import (
	"context"
	"server/internal/models/user"
)

type UserProfileService struct {
	userOps *user.Ops
}

func NewUserProfileService(userOps *user.Ops) *UserProfileService {
	return &UserProfileService{
		userOps: userOps,
	}
}

func (s *UserProfileService) GetUserProfile(ctx context.Context, userID uint) (*user.User, error) {
	return s.userOps.GetByID(ctx, userID)
}

func (s *UserProfileService) UpdateUserProfile(ctx context.Context, user *user.User) (*user.User, error) {
	return s.userOps.Update(ctx, user)
}
