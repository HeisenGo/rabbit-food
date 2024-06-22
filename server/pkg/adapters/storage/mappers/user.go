package mappers

import (
	"server/internal/models/user"
	"server/pkg/adapters/storage/entities"
)

func UserEntityToDomain(entity *entities.User) *user.User {
	return &user.User{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Email:     entity.Email,
		Password:  entity.Password,
		Phone:     entity.Phone,
		IsAdmin:  entity.IsAdmin,
	}
}

func UserDomainToEntity(domainUser *user.User) *entities.User {
	return &entities.User{
		Phone:    domainUser.Phone,
		Email:    domainUser.Email,
		Password: domainUser.Password,
	}
}
