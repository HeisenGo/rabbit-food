package models

type Address struct {
	UserID       uint        `json:"user_id"`
	AddressLine  string      `json:"address_line"`
	Coordinates  Coordinates `json:"coordinates"`
	Types        string      `json:"types"`
	City         string      `json:"city"`
	RestaurantID uint        `json:"restaurant_id"`
}

func NewAddress(addressLine string, coordinates Coordinates, types string, city string) *Address {
	return &Address{
		AddressLine: addressLine,
		Coordinates: coordinates,
		Types:       types,
		City:        city,
	}
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
