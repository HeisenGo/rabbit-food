package models

type RestaurantCategory struct {
	ID   uint   `json:"category_id"`
	Name string `json:"name"`
}

type RestaurantMenu struct {
	ID   uint   `json:"menu_id"`
	Name string `json:"name"`
}
