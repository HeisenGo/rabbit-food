package address

import (
	"context"
	"gorm.io/gorm"
)

type AddressOps struct {
	db   *gorm.DB
	repo Repo
}

func NewAddressOps(db *gorm.DB, repo Repo) *AddressOps {
	return &AddressOps{
		db:   db,
		repo: repo,
	}
}

func (o *AddressOps) Create(ctx context.Context, address *Address,userID uint) (*Address, error) {
	return o.repo.Create(ctx, address,userID)
}

func (o *AddressOps) GetByUser(ctx context.Context, userID uint) (*Address, error) {
	var address *Address
	address , err:= o.repo.GetByUser(ctx,userID)
	if err != nil{
		return nil,nil
	}
	return address, nil
}
func (o *AddressOps) GetByRestaurant(ctx context.Context, name string) (*Address, error) {
	var address *Address
	address , err:= o.repo.GetByRestaurant(ctx,name)
	if err != nil{
		return nil,nil
	}
	return address, nil

}
