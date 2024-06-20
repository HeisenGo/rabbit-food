package menus

import "bufio"

type MenuComponent interface {
	GetName() string
	Display()
	Execute(scanner *bufio.Scanner)
}
