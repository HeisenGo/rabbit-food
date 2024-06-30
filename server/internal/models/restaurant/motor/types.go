package motor

import (
	"context"
)

type Repo interface {
	Create(ctx context.Context, motor *Motor) (*Motor, error)
}

type Motor struct {
	ID           uint      `json:"id"`
	Name         string  `json:"name"`
	RestaurantID uint   `json:"restaurant_id"`
	Speed        int  `json:"speed"`
}

func NewMotor(restaurantID uint, speed int) *Motor {
	return &Motor{RestaurantID: restaurantID, Speed: speed}
}
