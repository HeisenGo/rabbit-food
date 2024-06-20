package main

import (
	"bufio"
	"client/config"
	"client/menus"
	"client/utils"
	"flag"
	"os"
)

func main() {
	utils.WelcomeSayer()
	defer utils.ByePrinter()
	cfg := readConfig()
	scanner := bufio.NewScanner(os.Stdin)

	mainMenu := menus.GetMainMenu(cfg.GetService())
	mainMenu.Execute(scanner)
}

var envFilePath = flag.String("envpath", "", ".env file path")

func readConfig() config.IConfig {
	flag.Parse()
	return config.ReadConfig(*envFilePath)
}
