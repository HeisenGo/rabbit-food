package menus

import "client/commands"

type IMenuItem interface {
	GetName() string
	GetCommand() commands.Command
	GetPostMenu() MenuComponent
}
