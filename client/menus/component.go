package menus

import "bufio"

type MenuComponent interface {
	Display()
	Execute(scanner *bufio.Scanner)
}
