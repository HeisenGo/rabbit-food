package address

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
)

type Repo interface {
	Create(ctx context.Context, address *Address) (*Address, error)
	Delete(ctx context.Context, addressID uint) error
	GetByID(ctx context.Context, addressID uint) (*Address, error)
	Update(ctx context.Context, address *Address) (*Address, error)
}

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

func (c Coordinates) GormDataType() string {
	return "geography(POINT, 4326)"
}

func (c Coordinates) Value() (driver.Value, error) {
	return fmt.Sprintf("SRID=4326;POINT(%f %f)", c.Lng, c.Lat), nil
}

func (c *Coordinates) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("type assertion to string failed")
	}
	_, err := fmt.Sscanf(str, "POINT(%f %f)", &c.Lng, &c.Lat)
	if err != nil {
		return err
	}
	return nil
}
