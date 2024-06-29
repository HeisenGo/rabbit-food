package motor

import (
	"context"
)

type Repo interface {
	Create(ctx context.Context, motor *Motor) (*Motor, error)
}

type Motor struct {
	ID           uint
	Name         string
	RestaurantID uint
	Speed        int
}

func NewMotor(restaurantID uint, speed int) *Motor {
	return &Motor{RestaurantID: restaurantID, Speed: speed}
}
