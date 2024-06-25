package address

import (
	"context"
	"server/internal/models/user"
)

type Repo interface {
	Create(ctx context.Context, address *Address) (*Address, error)
	GetByUser(ctx context.Context, userID uint) (*Address, error)
	GetByRestaurant(ctx context.Context, name string) (*Address, error)
}

type Address struct {
	UserID      uint 		
	Addressline string		
	Cordinates  string 		
	Types  		uint			
	City 		string 			
	
}
func NewAddress(addressline string, cordinates  string, types uint,city string ) *Address {
	return &Address{
		Addressline:    addressline,
		Cordinates:     cordinates,
		Types: 			types,
		City:			city,
	}
}

func (u *Address) SetAddressline(addressline string) {
	u.Addressline = addressline
}
func (u *Address) SetCordinates(cordinates string) {
	u.Cordinates = cordinates
}
func (u *Address) SetCity(city string) {
	u.City = city
}