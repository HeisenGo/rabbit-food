package services
import (
	"context"
	"server/internal/models/address"
)
type AddressService struct{
	addressOps *address.AddressOps
}
func NewAddressService(addressOps *address.AddressOps) *AddressService {
	return &AddressService{
		addressOps: addressOps,
	}
}
func (s *AddressService) Create(ctx context.Context, addressline string,cordinates [2]float64,types string,city string) (*address.Address, error) {
	
	newaddress := &address.Address{
		Addressline:addressline,
		Cordinates: cordinates,
		Types: types,
		City : city,
	}
	createdAddress, err := s.addressOps.Create(ctx, newaddress)
	if err != nil {
		return nil, err
	}
	return createdAddress, nil
}
