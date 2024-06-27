package mappers

import (
	"server/internal/models/restaurant/menu"
	"server/pkg/adapters/storage/entities"
)

func MenuEntityToDomain(entity *entities.Menu) *menu.Menu {
	return &menu.Menu{
		ID:           entity.ID,
		Name:         entity.Name,
		RestaurantID: entity.RestaurantID,
	}
}

func MenuItemEntityToDomain(entity *entities.MenuItem) *menu.MenuItem {
	return &menu.MenuItem{
		ID:              entity.ID,
		Name:            entity.Name,
		Price:           entity.Price,
		PreparationTime: entity.PreparationTime,
		MenuID:          entity.MenuID,
	}
}

func BatchMenuEntityToDomain(entities []*entities.Menu) []*menu.Menu {
	var domainMenus []*menu.Menu
	for _, e := range entities {
		domainMenus = append(domainMenus, &menu.Menu{ID: e.ID, Name: e.Name, RestaurantID: e.RestaurantID})
	}
	return domainMenus
}

func MenuDomainToEntity(domainMenu *menu.Menu) *entities.Menu {
	return &entities.Menu{
		Name:         domainMenu.Name,
		RestaurantID: domainMenu.RestaurantID,
	}
}

func MenuItemDomainToEntity(domainMenuItem *menu.MenuItem) *entities.MenuItem {
	return &entities.MenuItem{
		Name:            domainMenuItem.Name,
		Price:           domainMenuItem.Price,
		PreparationTime: domainMenuItem.PreparationTime,
		MenuID:          domainMenuItem.MenuID,
	}
}
