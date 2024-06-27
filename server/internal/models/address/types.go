package address

import (
	"context"
)

type Repo interface {
	Create(ctx context.Context, address *Address) (*Address, error)
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

