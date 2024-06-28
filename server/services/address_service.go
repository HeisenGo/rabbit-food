package services

import (
	"context"
	"server/internal/models/address"
)

type AddressService struct {
	addressOps *address.AddressOps
}

func NewAddressService(addressOps *address.AddressOps) *AddressService {
	return &AddressService{
		addressOps: addressOps,
	}
}
func (s *AddressService) Create(ctx context.Context, addressLine string, coordinates address.Coordinates, types string, city string) (*address.Address, error) {
	newAddress := &address.Address{
		AddressLine: addressLine,
		Coordinates: coordinates,
		Types:       types,
		City:        city,
	}
	createdAddress, err := s.addressOps.Create(ctx, newAddress)
	if err != nil {
		return nil, err
	}
	return createdAddress, nil
}
