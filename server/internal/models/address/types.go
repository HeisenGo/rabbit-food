package address

import (
	"context"
)

type Repo interface {
	Create(ctx context.Context, address *Address, userID uint) (*Address, error)
	GetByUser(ctx context.Context, userID uint) (*Address, error)
	GetByRestaurant(ctx context.Context, name string) (*Address, error)
}

type Address struct {
	UserID      uint 		
	Addressline string		
	Cordinates  [2]float64 		
	Types  		string			
	City 		string 			
	
}
func NewAddress(addressline string, cordinates  [2]float64, types string,city string ) *Address {
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
func (u *Address) SetCordinates(cordinates [2]float64) {
	u.Cordinates = cordinates
}
func (u *Address) SetCity(city string) {
	u.City = city
}
func (u *Address) SetUserAddress(types string){
	a := types==u.Types
	if a== true && (types == "User") {
		NewAddress(u.Addressline,u.Cordinates,u.Types,u.City)
	}  
}
func (u *Address) SetRestaurantAddress(types string){
	a := types==u.Types
	if a== true && (types == "Restaurant") {
		NewAddress(u.Addressline,u.Cordinates,u.Types,u.City)
	}  
}