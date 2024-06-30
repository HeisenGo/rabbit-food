package commands

import (
	"client/protocol/tcp"
	"client/services"
)

type AddItemToMenuCommand struct {
	service services.Service
}

func (c *AddItemToMenuCommand) Execute(addItemToMenuData tcp.AddItemToMenuReqBody) error {

	err := c.service.AddItemToMenu(&addItemToMenuData)
	if err != nil {
		return err
	}
	return nil
}

func NewAddItemToMenuCommand(service services.Service) *AddItemToMenuCommand {
	return &AddItemToMenuCommand{service: service}
}
