package utils

import (
	"bufio"
	"client/constants"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func ReadInput(scanner *bufio.Scanner, prompt string) string {
	ColoredPrint(constants.Purple, fmt.Sprintf("\t%s", prompt))
	scanner.Scan()
	return scanner.Text()
}

func ClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		fmt.Println("Unsupported platform")
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ColoredPrint(color constants.Color, a ...any) {
	var b []any
	b = append(b, color)
	b = append(b, a...)
	b = append(b, constants.Reset)
	fmt.Print(b...)
}

func ByePrinter() {
	ClearScreen()
	ColoredPrint(constants.Red, constants.AppLogo)
}

func WelcomeSayer() {
	ClearScreen()
	ColoredPrint(constants.Red, constants.AppLogo)
	time.Sleep(2 * time.Second)
	ClearScreen()
}
