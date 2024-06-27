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
func (s *AddressService) Create(ctx context.Context, addressline string,cordinates [2]float64,types string,city string,userID uint) (*address.Address, error) {
	
	newaddress := &address.Address{
		Addressline:addressline,
		Cordinates: cordinates,
		Types: types,
		City : city,
	}
	createdAddress, err := s.addressOps.Create(ctx, newaddress,userID)
	if err != nil {
		return nil, err
	}
	return createdAddress, nil
}
func (s *AddressService) GetByUser(ctx context.Context,userID uint ,types uint) (*address.Address, error) {
	//TODO : define the number for types to get the address from users table or restaurant
	useraddress , err := s.addressOps.GetByUser(ctx,userID)
	if err !=nil {
		return nil,err
	}
	return useraddress,nil
}
func (s *AddressService) GetByRestaurant(ctx context.Context,name string ,types uint) (*address.Address, error) {
	//TODO : define the number for types to get the address from users table or restaurant
	restaurantaddress , err := s.addressOps.GetByRestaurant(ctx,name)
	if err !=nil {
		return nil,err
	}
	return restaurantaddress,nil
}