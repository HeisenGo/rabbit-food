package services

import (
	"context"
	"server/internal/models/address"
	"server/internal/models/user"
	"server/pkg/utils/users"
)

type UserService struct {
	userOps    *user.Ops
	addressOps *address.AddressOps
}

func NewUserService(userOps *user.Ops, addressOps *address.AddressOps) *UserService {
	return &UserService{
		userOps:    userOps,
		addressOps: addressOps,
	}
}

func (s *UserService) GetUserByEmailOrPhone(ctx context.Context, phoneOrEmail string) (*user.User, error) {
	user, err := s.userOps.GetOperatorUser(ctx, phoneOrEmail)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*user.User, error) {
	return s.userOps.GetByID(ctx, id)
}

func (s *UserService) UpdateUserFirstName(ctx context.Context, id uint, firstName string) (*user.User, error) {
	return s.userOps.UpdateField(ctx, id, "first_name", firstName)
}

func (s *UserService) UpdateUserLastName(ctx context.Context, id uint, lastName string) (*user.User, error) {
	return s.userOps.UpdateField(ctx, id, "last_name", lastName)
}

func (s *UserService) UpdateUserEmail(ctx context.Context, id uint, email string) (*user.User, error) {
	return s.userOps.UpdateField(ctx, id, "email", email)
}

func (s *UserService) UpdateUserPhone(ctx context.Context, id uint, phone string) (*user.User, error) {
	return s.userOps.UpdateField(ctx, id, "phone", phone)
}

func (s *UserService) UpdateUserPassword(ctx context.Context, id uint, password string) (*user.User, error) {
	hashedPass, err := users.HashPassword(password)
	if err != nil {
		return nil, err
	}
	return s.userOps.UpdateField(ctx, id, "password", hashedPass)
}

func (s *UserService) DeleteUserAddress(ctx context.Context, userID, addressID uint) error {
	return s.addressOps.Delete(ctx, addressID)
}

func (s *UserService) AddUserAddress(ctx context.Context, userID uint, addr *address.Address) (*address.Address, error) {
	addr.UserID = userID
	return s.addressOps.Create(ctx, addr)
}
